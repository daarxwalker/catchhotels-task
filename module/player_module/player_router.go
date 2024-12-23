package player_module

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/middleware"
	"catchhotels/module/player_module/auth_handler"
	"catchhotels/module/player_module/character_handler"
)

func Register(router fiber.Router) {
	{
		auth := router.Group("/auth")
		auth.Post("/register", auth_handler.Register())
		auth.Post("/login", auth_handler.Login())
		auth.Delete("/logout", auth_handler.Logout())
	}
	{
		auth := router.Group("/character").Use(middleware.AccessGuard())
		auth.Post("/create", character_handler.CreateCharacter())
		auth.Post("/use/:characterId", character_handler.UseCharacter())
	}
}
