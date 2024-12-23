package dungeon_master_module

import (
	"github.com/gofiber/fiber/v2"

	"catchhotels/module/dungeon_master_module/player_handler"
)

func Register(router fiber.Router) {
	{
		player := router.Group("/player")
		player.Get("", player_handler.FindPlayer())
	}
}
