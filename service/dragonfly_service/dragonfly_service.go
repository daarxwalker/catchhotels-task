package dragonfly_service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	"catchhotels/config/dragonfly_config"
)

type DragonflyService struct {
	client *redis.Client
}

const (
	Token = "dragonfly_service"
)

func New(cfg *viper.Viper) *DragonflyService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := redis.NewClient(
		&redis.Options{
			Addr:     cfg.GetString(dragonfly_config.Uri),
			Password: cfg.GetString(dragonfly_config.Password),
			DB:       cfg.GetInt(dragonfly_config.DB),
		},
	)
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Ping to DragonflyDB failed: %v", err)
	}
	fmt.Println("Successfully connected and pinged DragonflyDB!")
	return &DragonflyService{
		client: client,
	}
}

func (s *DragonflyService) Exists(c *fiber.Ctx, key ...string) bool {
	return s.client.Exists(c.Context(), key...).Val() == 1
}

func (s *DragonflyService) Get(c *fiber.Ctx, key string, target any) error {
	value := s.client.Get(c.Context(), key).Val()
	if len(value) == 0 {
		return nil
	}
	return json.Unmarshal([]byte(value), target)
}

func (s *DragonflyService) MustGet(c *fiber.Ctx, key string, target any) {
	if err := s.Get(c, key, target); err != nil {
		panic(err)
	}
}

func (s *DragonflyService) Set(c *fiber.Ctx, key string, value any, expiration time.Duration) error {
	valueBytes, marshalErr := json.Marshal(value)
	if marshalErr != nil {
		return marshalErr
	}
	return s.client.Set(c.Context(), key, string(valueBytes), expiration).Err()
}

func (s *DragonflyService) MustSet(c *fiber.Ctx, key string, value any, expiration time.Duration) {
	if err := s.Set(c, key, value, expiration); err != nil {
		panic(err)
	}
}

func (s *DragonflyService) Destroy(c *fiber.Ctx, key string) error {
	return s.Set(c, key, "", time.Millisecond)
}

func (s *DragonflyService) MustDestroy(c *fiber.Ctx, key string) {
	if err := s.Destroy(c, key); err != nil {
		panic(err)
	}
}
