package auction_websockets

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/middlewares"
	auctionReq "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strconv"

	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/redisdb"
)

func Auction(c *websocket.Conn) {

	// Get depends, create context with dependency
	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	// Set depends in context
	ctx, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
		}
		return
	}

	authReq, errResponse := middlewares.AuthWS(c, ctx)
	if errResponse != nil {
		if err = c.WriteJSON(err); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
		}

		if err = c.Close(); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to close websocket connection with error: %s", err.Error())))
		}
		return
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
		}
		return
	}

	context_with_depends.StartDBTx(ctx)
	defer context_with_depends.DBTxRollback(ctx)

	auction, err := db_repository.AuctionInterface.Auction(ctx, id)
	if err != nil || auction == nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
		}
		return
	}

	if auction.Status != entities.ActiveStatus {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
		}
		return
	}

	for {
		var member *entities.AuctionMember
		auctionOffer := &auctionReq.AmountOffer{}

		if err = c.ReadJSON(auctionOffer); err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to parser request"}); err != nil {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
			}
			continue
		}

		member, err = db_repository.AuctionInterface.Member(ctx, auction.ID, authReq.ID)
		if err != nil || member == nil || len(member.ParticipantID) <= 0 {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
			}
			continue
		}

		if auctionOffer.Amount <= auction.MinPrice {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "your amount lower than minimal price"}); err != nil {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
			}
			continue
		}

		if auctionOffer.Amount <= auction.Price {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "your amount lower than current price"}); err != nil {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
			}
			continue
		}

		if err = db_repository.AuctionInterface.UpdatePrice(ctx, auction.ID, auctionOffer.Amount); err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to set price for some reason"}); err != nil {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
			}
			continue
		}

		context_with_depends.DBTxCommit(ctx)

		if err = c.WriteJSON(fiber.Map{"status": "success", "price": auctionOffer.Amount}); err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to set price for some reason"}); err != nil {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to write json in websocket with error: %s", err.Error())))
			}
			continue
		}
	}
}
