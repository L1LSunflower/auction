package app

import (
	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/handlers"
	"github.com/L1LSunflower/auction/internal/workers"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"strconv"
)

func App() {
	// App with own config
	app := fiber.New(fiber.Config{
		BodyLimit:               100 * 1024 * 1024,
		ServerHeader:            "accept",
		StrictRouting:           true,
		EnableTrustedProxyCheck: true,
		RequestMethods:          []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	})
	// Cors
	app.Use(cors.New())
	// internal Config
	cfg := config.GetConfig()
	logger.Log = logger.New(cfg.Log.Level, cfg.Log.Driver)
	// Set routes
	handlers.SetRoutes(app)
	// Start workers
	workers.StartWorkers()
	// App Listener
	if err := app.Listen("0.0.0.0:" + strconv.Itoa(cfg.AppPort)); err != nil {
		log.Fatal(err)
	}
}
