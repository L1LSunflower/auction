package auction_websockets

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	auctionService "github.com/L1LSunflower/auction/internal/domain/services/auctions"
	"github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
	"github.com/L1LSunflower/auction/pkg/redisdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strconv"
	"time"
)

func Auction(c *websocket.Conn) {
	var (
		auctionID int
		wsID      int
		err       error
	)

	bearerToken := config.GetConfig().BearerToken
	auth := &auctions.WSAuth{}

	if err = c.ReadJSON(auth); err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to parse request"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
		}
		return
	}

	if bearerToken != auth.Bearer {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "auth required"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
		}
		return
	}

	auctionID, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction id"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
		}
		return
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	ctx, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get dependencies"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
		}
		return
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, auctionID)
	if err != nil || auction == nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction id"}); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
		}
		return
	}

	if err = c.WriteJSON(fiber.Map{"status": "success", "price": auction.Price}); err != nil {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
	}

	wsID = auctionService.LengthStackOfAuction(auctionID)
	if wsID == 0 {
		wsID = 1
	}
	auctionService.RegisterNew(auctionID, wsID)

	for {
		Update(c, auctionID, wsID)
	}

}

func Update(c *websocket.Conn, auctionID, wsID int) {
	if !auctionService.CheckEvent(auctionID) {
		time.Sleep(5 * time.Second)
		Update(c, auctionID, wsID)
	}

	if auctionService.CheckActual(auctionID, wsID) {
		time.Sleep(1 * time.Second)
		Update(c, auctionID, wsID)
	}

	price := auctionService.AuctionPrice(auctionID)

	if err := c.WriteJSON(fiber.Map{"status": "success", "price": price}); err != nil {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
	}

	auctionService.SetActual(auctionID, wsID)
}
