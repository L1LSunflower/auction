package handlers

import (
	"github.com/gofiber/fiber/v2"

	cityHandler "github.com/L1LSunflower/auction/internal/handlers/cities"
	usersHandler "github.com/L1LSunflower/auction/internal/handlers/users"
	"github.com/L1LSunflower/auction/internal/middlewares"
	cityValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/cities"
	usersValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/users"
)

func SetRoutes(app *fiber.App) {
	app.Get("/__health", Healthcheck)

	v1 := app.Group("/v1")
	v1.Use(middlewares.Attempts())
	v1.Use(middlewares.BearerAuth())

	// Cities routes
	v1.Post("/cities", cityValidator.Create, cityHandler.Create)
	v1.Get("/cities", cityHandler.Cities)
	v1.Get("/cities/:id", cityValidator.City, cityHandler.City)
	v1.Put("/cities/:id", cityValidator.UpdateCity, cityHandler.Update)
	v1.Delete("/cities/:id", cityValidator.DeleteCity, cityHandler.Delete)

	// Auth routes
	v1.Post("/sign-up", usersValidator.SignUpValidator, usersHandler.SignUp)
	v1.Post("/sign-in", usersValidator.SignInValidator, usersHandler.SignIn)
	v1.Post("/confirm/:id", usersValidator.ConfirmValidator, usersHandler.Confirm)
	//v1.Post("/restore-password", users.RestorePassword)
	//v1.Get("/refresh_token", middlewares.usersAuth, usersHandler.Refresh)
	v1.Get("/user/:id", usersValidator.GetUserValidator, usersHandler.GetUser)

	// Auction
	//v1.Post("/auctions", auction.NewAuction)
	//v1.Get("/auctions", auction.Auctions)
	//v1.Get("/auctions/:id", auction.Auction)
	//v1.Put("/auctions/:id", auction.UpdateAuction)
	//v1.Delete("/auctions/id", auction.DeleteAuction)
}
