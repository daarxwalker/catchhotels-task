package bootstrap

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	
	"catchhotels/config"
	"catchhotels/config/app_config"
	"catchhotels/middleware"
	"catchhotels/module/dungeon_master_module"
	"catchhotels/module/player_module"
)

func Boot() error {
	cfg := config.Read()
	app := fiber.New(
		fiber.Config{
			AppName:      cfg.GetString(app_config.Name),
			ErrorHandler: middleware.ErrorHandler(),
		},
	)
	
	// Swagger
	app.Use(
		swagger.New(
			swagger.Config{
				BasePath: "/",
				FilePath: "./docs/swagger.json",
				Path:     "docs",
				Title:    "Catchhotels Task Documentation",
			},
		),
	)
	
	// Middleware
	app.Use(middleware.ServiceProvider(cfg))
	app.Use(middleware.ResponseInterceptor())
	
	// Module
	dungeon_master_module.Register(
		app.Group("/dungeon-master").Use(middleware.DungeonMaster(cfg)),
	)
	player_module.Register(app.Group("/player"))
	
	return app.Listen(":" + cfg.GetString(app_config.Port))
}
