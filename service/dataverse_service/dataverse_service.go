package dataverse_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	
	"catchhotels/config/dataverse_config"
)

type DataverseService struct {
	cfg         *viper.Viper
	bearerToken string
}

const (
	Token = "dataverse_service"
)

func New(cfg *viper.Viper) (*DataverseService, error) {
	instance := &DataverseService{
		cfg: cfg,
	}
	go func() {
		for {
			bearerToken, authorizeErr := instance.Authorize()
			if authorizeErr != nil {
				panic(authorizeErr)
			}
			instance.bearerToken = bearerToken
			time.Sleep(cfg.GetDuration(dataverse_config.AuthRenewInterval))
		}
	}()
	return instance, nil
}

func (s *DataverseService) Authorize() (string, error) {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", s.cfg.GetString(dataverse_config.ClientId))
	formData.Set("client_secret", s.cfg.GetString(dataverse_config.ClientSecret))
	formData.Set("scope", s.cfg.GetString(dataverse_config.AuthScope))
	res, postErr := http.Post(
		s.cfg.GetString(dataverse_config.AuthEndpoint),
		fiber.MIMEApplicationForm,
		bytes.NewBufferString(formData.Encode()),
	)
	if postErr != nil {
		return "", postErr
	}
	defer func() {
		if closeBodyErr := res.Body.Close(); closeBodyErr != nil {
			panic(closeBodyErr)
		}
	}()
	var authRes AuthResponse
	if decodeErr := json.NewDecoder(res.Body).Decode(&authRes); decodeErr != nil {
		return "", decodeErr
	}
	return authRes.AccessToken, nil
}

func (s *DataverseService) Find(tableName string, target any) error {
	if bearerTokenErr := s.validateBearerToken(); bearerTokenErr != nil {
		return bearerTokenErr
	}
	req, createReqErr := http.NewRequest(http.MethodGet, s.cfg.GetString(dataverse_config.Endpoint)+"/"+tableName, nil)
	if createReqErr != nil {
		return createReqErr
	}
	if sendReqErr := s.sendRequest(req, target); sendReqErr != nil {
		return sendReqErr
	}
	return nil
}

func (s *DataverseService) MustFind(tableName string, target any) {
	if err := s.Find(tableName, target); err != nil {
		panic(err)
	}
}

func (s *DataverseService) FindOne(tableName, id string, target any) error {
	if bearerTokenErr := s.validateBearerToken(); bearerTokenErr != nil {
		return bearerTokenErr
	}
	req, createReqErr := http.NewRequest(
		http.MethodGet, s.cfg.GetString(dataverse_config.Endpoint)+"/"+tableName+"("+id+")", nil,
	)
	if createReqErr != nil {
		return createReqErr
	}
	if sendReqErr := s.sendRequest(req, target); sendReqErr != nil {
		return sendReqErr
	}
	return nil
}

func (s *DataverseService) MustFindOne(tableName, id string, target any) {
	if err := s.Find(tableName, target); err != nil {
		panic(err)
	}
}

func (s *DataverseService) FindWithEmail(tableName, email string, target any) error {
	if bearerTokenErr := s.validateBearerToken(); bearerTokenErr != nil {
		return bearerTokenErr
	}
	params := url.Values{}
	params.Set("$filter", fmt.Sprintf("cr568_email eq '%s'", email))
	params.Set("$top", "1")
	endpoint := s.cfg.GetString(dataverse_config.Endpoint) + "/" + tableName + "?" + params.Encode()
	req, createReqErr := http.NewRequest(http.MethodGet, endpoint, nil)
	if createReqErr != nil {
		return createReqErr
	}
	if sendReqErr := s.sendRequest(req, target); sendReqErr != nil {
		return sendReqErr
	}
	return nil
}

func (s *DataverseService) MustFindWithEmail(tableName, email string, target any) {
	if err := s.FindWithEmail(tableName, email, target); err != nil {
		panic(err)
	}
}

func (s *DataverseService) Create(tableName string, data, target any) error {
	if bearerTokenErr := s.validateBearerToken(); bearerTokenErr != nil {
		return bearerTokenErr
	}
	bodyBytes, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		return marshalErr
	}
	req, createReqErr := http.NewRequest(
		http.MethodPost,
		s.cfg.GetString(dataverse_config.Endpoint)+"/"+tableName,
		bytes.NewReader(bodyBytes),
	)
	if createReqErr != nil {
		return createReqErr
	}
	req.Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	if sendReqErr := s.sendRequest(req, target); sendReqErr != nil {
		return sendReqErr
	}
	return nil
}

func (s *DataverseService) MustCreate(tableName string, data, target any) {
	if err := s.Create(tableName, data, target); err != nil {
		panic(err)
	}
}

func (s *DataverseService) validateBearerToken() error {
	if len(s.bearerToken) == 0 {
		return errors.New("invalid bearer token")
	}
	return nil
}

func (s *DataverseService) sendRequest(req *http.Request, target any) error {
	req.Header.Add("Authorization", s.bearerToken)
	client := &http.Client{}
	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr
	}
	defer func() {
		if closeBodyErr := res.Body.Close(); closeBodyErr != nil {
			panic(closeBodyErr)
		}
	}()
	if target != nil && res.ContentLength > 0 {
		if decodeErr := json.NewDecoder(res.Body).Decode(target); decodeErr != nil {
			return decodeErr
		}
	}
	return nil
}

func (s *DataverseService) createQuery(values map[string]string) string {
	var query []string
	for key, value := range values {
		query = append(query, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(query, "&")
}
