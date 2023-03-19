package handlers

import (
	auctionHandler "github.com/L1LSunflower/auction/internal/handlers/auction"
	"github.com/L1LSunflower/auction/internal/handlers/auction_websockets"
	balanceHandler "github.com/L1LSunflower/auction/internal/handlers/balances"
	"github.com/L1LSunflower/auction/internal/handlers/fileuploader"
	tagsHandler "github.com/L1LSunflower/auction/internal/handlers/tags"
	usersHandler "github.com/L1LSunflower/auction/internal/handlers/users"
	"github.com/L1LSunflower/auction/internal/middlewares"
	auctionValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/auctions"
	balanceValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/balances"
	tagsValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/tags"
	usersValidator "github.com/L1LSunflower/auction/internal/middlewares/validator/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/__health", Healthcheck)

	app.Get("/auctions/:id", middlewares.Attempts(), websocket.New(auction_websockets.Auction))

	v1 := app.Group("/v1")
	v1.Use(middlewares.Attempts())
	v1.Use(middlewares.BearerAuth())

	// Auth routes
	v1.Post("/sign-up", usersValidator.SignUpValidator, usersHandler.SignUp)
	v1.Post("/sign-in", usersValidator.SignInValidator, usersHandler.SignIn)
	v1.Post("/confirm/:id", usersValidator.ConfirmValidator, usersHandler.Confirm)
	v1.Post("/restore_password", usersValidator.RestoreValidator, usersHandler.Restore)
	v1.Put("/restore_password", usersValidator.ChangePasswordValidator, usersHandler.ChangePassword)
	v1.Post("/refresh_token", middlewares.Auth(), usersValidator.RefreshValidator, usersHandler.Refresh)

	// Users routes
	v1.Get("/profile", middlewares.Auth(), usersValidator.ProfileValidator, usersHandler.Profile)
	v1.Get("/profile_history", middlewares.Auth(), usersValidator.ProfileValidator, usersHandler.ProfileHistory)
	v1.Get("/users", middlewares.Auth(), usersValidator.GetUserValidator, usersHandler.GetUser)
	v1.Put("/users", middlewares.Auth(), usersValidator.UpdateValidator, usersHandler.Update)
	v1.Delete("/users", middlewares.Auth(), usersValidator.DeleteValidator, usersHandler.Delete)

	// Auction
	v1.Post("/auctions", middlewares.Auth(), auctionValidator.Create, auctionHandler.Create)
	v1.Get("/auctions", middlewares.Auth(), auctionValidator.Auctions, auctionHandler.Auctions)
	v1.Get("/auctions/:id", middlewares.Auth(), auctionValidator.Auction, auctionHandler.Auction)
	v1.Put("/auctions/:id", middlewares.Auth(), auctionValidator.Update, auctionHandler.Update)
	v1.Delete("/auctions/:id", middlewares.Auth(), auctionValidator.Delete, auctionHandler.Delete)
	v1.Post("/auctions/:id/start", middlewares.Auth(), auctionValidator.Start, auctionHandler.Start)
	v1.Post("/auctions/:id/end", middlewares.Auth(), auctionValidator.End, auctionHandler.End)
	v1.Post("/auctions/:id/participate", middlewares.Auth(), auctionValidator.Participate, auctionHandler.Participate)

	// Balance routes
	balance := v1.Group("/balance")
	balance.Post("/credit", middlewares.Auth(), balanceValidator.Credit, balanceHandler.Credit)
	balance.Post("/debit", middlewares.Auth(), balanceValidator.Debit, balanceHandler.Debit)
	balance.Get("/get_balance", middlewares.Auth(), balanceValidator.Balance, balanceHandler.Balance)

	// Get tags like
	v1.Get("/tags", tagsValidator.ByPattern, tagsHandler.ByPattern)

	// File uploader
	app.Static("/static", "./static")
	app.Post("/upload_file", middlewares.BearerAuth(), fileuploader.UploadFile)
	app.Delete("/delete_file/:fileName", middlewares.BearerAuth(), fileuploader.DeleteFile)
}
