package handlers

import (
	"github.com/L1LSunflower/auction/internal/handlers/fileuploader"
	"github.com/gofiber/fiber/v2"

	auctionHandler "github.com/L1LSunflower/auction/internal/handlers/auction"
	usersHandler "github.com/L1LSunflower/auction/internal/handlers/users"
	"github.com/L1LSunflower/auction/internal/middlewares"
	auctionValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/auctions"
	usersValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/users"
)

func SetRoutes(app *fiber.App) {
	app.Get("/__health", Healthcheck)

	v1 := app.Group("/v1")
	v1.Use(middlewares.Attempts())
	v1.Use(middlewares.BearerAuth())

	// Auth routes
	v1.Post("/sign-up", usersValidator.SignUpValidator, usersHandler.SignUp)
	v1.Post("/sign-in", usersValidator.SignInValidator, usersHandler.SignIn)
	v1.Post("/confirm/:id", usersValidator.ConfirmValidator, usersHandler.Confirm)
	v1.Post("/restore_password", usersValidator.RestoreValidator, usersHandler.Restore)
	v1.Get("/refresh_token", middlewares.Auth(), usersValidator.RefreshValidator, usersHandler.Refresh)
	v1.Get("/user/:id", middlewares.Auth(), usersValidator.GetUserValidator, usersHandler.GetUser)

	// Auction
	v1.Post("/auctions", middlewares.Auth(), auctionValidator.Create, auctionHandler.Create)
	v1.Get("/auctions", middlewares.Auth(), auctionValidator.Auctions, auctionHandler.Auctions)
	v1.Get("/auctions/:id", middlewares.Auth(), auctionValidator.Auction, auctionHandler.Auction)
	v1.Put("/auctions/:id", middlewares.Auth(), auctionValidator.Update, auctionHandler.Update)
	v1.Delete("/auctions/:id", middlewares.Auth(), auctionValidator.Delete, auctionHandler.Delete)
	v1.Post("/auctions/:id/start", middlewares.Auth(), auctionValidator.Start, auctionHandler.Start)
	v1.Post("/auctions/:id/end", middlewares.Auth(), auctionValidator.End, auctionHandler.End)

	// File uploader
	v1.Post("/upload_file", fileuploader.UploadFile)
	v1.Delete("/delete_file/:fileName", fileuploader.DeleteFile)
}
