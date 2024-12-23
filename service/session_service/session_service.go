package session_service

import (
	"github.com/dchest/uniuri"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"catchhotels/config/app_config"
	"catchhotels/config/auth_config"
	"catchhotels/service/dragonfly_service"
)

type SessionService struct {
	cfg              *viper.Viper
	dragonflyService *dragonfly_service.DragonflyService
}

const (
	Token = "session_service"
)

const (
	CookieKey = "X-Session"
)

func New(cfg *viper.Viper, dragonflyService *dragonfly_service.DragonflyService) *SessionService {
	return &SessionService{
		cfg:              cfg,
		dragonflyService: dragonflyService,
	}
}

func (s *SessionService) Get(c *fiber.Ctx) (Session, error) {
	sessionToken := c.Cookies(CookieKey)
	if len(sessionToken) == 0 {
		return Session{}, nil
	}
	var session Session
	if getSessionErr := s.dragonflyService.Get(c, Token+":"+sessionToken, &session); getSessionErr != nil {
		return session, getSessionErr
	}
	return session, nil
}

func (s *SessionService) MustGet(c *fiber.Ctx) Session {
	session, getErr := s.Get(c)
	if getErr != nil {
		panic(getErr)
	}
	return session
}

func (s *SessionService) Set(c *fiber.Ctx, session Session) error {
	var token string
	sessionToken := c.Cookies(CookieKey)
	if len(sessionToken) > 0 {
		token = sessionToken
	}
	if len(token) == 0 {
		token = uniuri.New()
	}
	expiration := s.cfg.GetDuration(auth_config.SessionExpiration)
	c.Cookie(
		&fiber.Cookie{
			Name:   CookieKey,
			Value:  token,
			Path:   "/",
			MaxAge: int(expiration.Seconds()),
			Secure: s.cfg.GetString(app_config.Env) == app_config.EnvDevelopment,
		},
	)
	return s.dragonflyService.Set(c, Token+":"+token, session, expiration)
}

func (s *SessionService) MustSet(c *fiber.Ctx, session Session) {
	if err := s.Set(c, session); err != nil {
		panic(err)
	}
}

func (s *SessionService) Renew(c *fiber.Ctx) error {
	sessionToken := c.Cookies(CookieKey)
	if len(sessionToken) == 0 {
		return nil
	}
	session, err := s.Get(c)
	if err != nil {
		return err
	}
	expiration := s.cfg.GetDuration(auth_config.SessionExpiration)
	c.Cookie(
		&fiber.Cookie{
			Name:   CookieKey,
			Value:  sessionToken,
			Path:   "/",
			MaxAge: int(expiration.Seconds()),
			Secure: s.cfg.GetString(app_config.Env) == app_config.EnvDevelopment,
		},
	)
	return s.dragonflyService.Set(c, Token+":"+sessionToken, session, expiration)
}

func (s *SessionService) MustRenew(c *fiber.Ctx) {
	if err := s.Renew(c); err != nil {
		panic(err)
	}
}

func (s *SessionService) Destroy(c *fiber.Ctx) error {
	sessionToken := c.Cookies(CookieKey)
	if len(sessionToken) == 0 {
		return nil
	}
	c.Cookie(
		&fiber.Cookie{
			Name:   CookieKey,
			Path:   "/",
			MaxAge: -1,
			Secure: s.cfg.GetString(app_config.Env) == app_config.EnvDevelopment,
		},
	)
	if len(sessionToken) > 0 {
		return s.dragonflyService.Destroy(c, Token+":"+sessionToken)
	}
	return nil
}

func (s *SessionService) MustDestroy(c *fiber.Ctx) {
	if err := s.Destroy(c); err != nil {
		panic(err)
	}
}
