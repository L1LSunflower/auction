package auction_websockets

import (
	"fmt"
	"github.com/L1LSunflower/auction/config"
	auctionService "github.com/L1LSunflower/auction/internal/domain/services/auctions"
	"github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
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
	auth := auctions.WSAuth{}

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
