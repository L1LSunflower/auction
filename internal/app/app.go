package app

import (
	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/handlers"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"strconv"
)

func App() {
	app := fiber.New(fiber.Config{
		//Prefork:                 true,
		ServerHeader:            "accept",
		StrictRouting:           true,
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"http://localhost:3000", "http://localhost:8100", "http://192.168.0.16:3000", "http://192.168.0.14:8100"},
		RequestMethods:          []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	})

	app.Use(cors.New())

	cfg := config.GetConfig()
	logger.Log = logger.New(cfg.Log.Level, cfg.Log.Driver, "")

	handlers.SetRoutes(app)

	if err := app.Listen("0.0.0.0:" + strconv.Itoa(cfg.AppPort)); err != nil {
		log.Fatal(err)
	}

	//if err := sms.SendSMS("+77479824031", "test code: 1234"); err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("success send sms!")
}
