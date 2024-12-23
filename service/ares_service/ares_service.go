package ares_service

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"

	"catchhotels/config/ares_config"
)

type AresService struct {
	endpoint string
}

const (
	Token = "ares_service"
)

func New(cfg *viper.Viper) *AresService {
	return &AresService{
		endpoint: cfg.GetString(ares_config.Endpoint),
	}
}

func (s *AresService) Get(cin string) (Response, error) {
	res, getErr := http.Get(s.endpoint + cin)
	if getErr != nil {
		return Response{}, getErr
	}
	defer func() {
		if closeBodyErr := res.Body.Close(); closeBodyErr != nil {
			panic(closeBodyErr)
		}
	}()
	var aresResult Response
	if decodeErr := json.NewDecoder(res.Body).Decode(&aresResult); decodeErr != nil {
		return Response{}, decodeErr
	}
	return aresResult, nil
}

func (s *AresService) MustGet(cin string) Response {
	res, err := s.Get(cin)
	if err != nil {
		panic(err)
	}
	return res
}
