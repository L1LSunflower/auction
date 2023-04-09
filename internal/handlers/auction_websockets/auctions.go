package auction_websockets

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	auctionService "github.com/L1LSunflower/auction/internal/domain/services/auctions"
	"github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
	"github.com/L1LSunflower/auction/pkg/redisdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"strconv"
	"time"
)

func Auction(c *websocket.Conn) {
	var (
		auctionID int
		err       error
	)

	//defer auctionService.DeleteConsumer(auctionID, c)

	bearerToken := config.GetConfig().BearerToken
	auth := &auctions.WSAuth{}

	log.Printf("\nSTEP: 1: SET CONNECTION\n\n")

	if err = c.ReadJSON(auth); err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to parse request"}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}
		return
	}

	log.Printf("\nSTEP: 2: CHECK BEARER\n\n")
	if bearerToken != auth.Bearer {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "auth required"}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}
		return
	}

	auctionID, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction id"}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}
		return
	}

	log.Printf("\nSTEP: 3: GET AUCTION ID\n\n")

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	log.Printf("\nSTEP: 4: SET DEPENDENCY\n\n")
	ctx, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get dependencies"}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}
		return
	}

	log.Printf("\nSTEP: 5: GET AUCTION BY ID\n\n")
	auction, err := db_repository.AuctionInterface.Auction(ctx, auctionID)
	if err != nil || auction == nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction id"}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}
		return
	}

	log.Printf("\nSTEP: 6: CHECK ON ONWER\n\n")
	if auth.UserID != auction.OwnerID {
		log.Printf("\nSTEP: 7: SEND DATA\n\n")
		if err = c.WriteJSON(fiber.Map{"status": "success", "user_id": auction.WinnerID, "price": auction.Price}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}

		log.Printf("\nSTEP: 7: REGISTER NEW IN STACK\n\n")
		auctionService.RegisterNew(auctionID, c)
		log.Printf("\nSTEP: 8: SET DELETER\n\n")
		go Delete(c, auctionID)
		for {
			UpdateParticipant(c, ctx, auctionID, auction.WinnerID)
			time.Sleep(1 * time.Second)
		}

	} else {
		log.Printf("\nSTEP: 8: GET MEMBERS\n\n")
		members, err := db_repository.AuctionInterface.Members(ctx, auctionID)
		if err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get owner"}); err != nil && err == c.Close() {
				logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
				auctionService.DeleteConsumer(auctionID, c)
			}
			return
		}

		log.Printf("\nSTEP: 9: PREPARE AUCTION MEMBERS\n\n")
		response := responses.AuctionMembers(members, fiber.Map{"user_id": auction.WinnerID, "price": auction.Price})
		if err = c.WriteJSON(response); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
			return
		}

		log.Printf("\nSTEP: 10: REGISTRE NEW IN STACK\n\n")
		auctionService.RegisterNew(auctionID, c)
		log.Printf("\nSTEP: 11: SET DELETER\n\n")
		go Delete(c, auctionID)
		for {
			UpdateOwner(c, ctx, auctionID, auction.WinnerID)
			time.Sleep(1 * time.Second)
		}
	}
}

func Delete(c *websocket.Conn, auctionID int) {
	closeReq := &auctions.WSClose{}
	if err := c.ReadJSON(closeReq); err != nil && err == c.Close() {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to read data from ws with error: %s", err)))
		auctionService.DeleteConsumer(auctionID, c)
	}

	if closeReq.Close {
		if err := c.Close(); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to close connection with error: %s", err)))
			auctionService.DeleteConsumer(auctionID, c)
		}
	}

	auctionService.DeleteConsumer(auctionID, c)
}

func UpdateParticipant(c *websocket.Conn, ctx context.Context, auctionID int, userID string) {
	if !auctionService.CheckEvent(auctionID) {
		//UpdateParticipant(c, ctx, auctionID, userID)
		return
	}

	if auctionService.CheckActual(auctionID, c) {
		//UpdateParticipant(c, ctx, auctionID, userID)
		return
	}

	log.Printf("\nSTEP: 9: GET AUCTION PRICE\n\n")
	userAndPrice := auctionService.AuctionPrice(auctionID)

	log.Printf("\nSTEP: 10: SEND AUCTION PARTICIPANT PRICE DATE\n\n")
	log.Printf("\nSTEP: 10.1: USER AND PRICE DATA: %#v\n\n", userAndPrice)
	if err := c.WriteJSON(fiber.Map{"status": "success", "user_id": userAndPrice["user_id"], "price": userAndPrice["price"]}); err != nil && err == c.Close() {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
		// Delete consumer
		auctionService.DeleteConsumer(auctionID, c)
		return
	}

	log.Printf("\nSTEP: 11: SET ACTUAL AUCTION\n\n")
	auctionService.SetActual(auctionID, c)

	//UpdateParticipant(c, ctx, auctionID, userID)
	return
}

func UpdateOwner(c *websocket.Conn, ctx context.Context, auctionID int, userID string) {
	if !auctionService.CheckEvent(auctionID) {
		//UpdateOwner(c, ctx, auctionID, userID)
		return
	}

	if auctionService.CheckActual(auctionID, c) {

		//UpdateOwner(c, ctx, auctionID, userID)
		return
	}

	log.Printf("\nSTEP: 12: GET AUCTION PRICE\n\n")
	userAndPrice := auctionService.AuctionPrice(auctionID)

	log.Printf("\nSTEP: 13: GET MEMBERS\n\n")
	members, err := db_repository.AuctionInterface.Members(ctx, auctionID)
	if err != nil {
		if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get members"}); err != nil && err == c.Close() {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))
			// Delete consumer
			auctionService.DeleteConsumer(auctionID, c)
			return
		}
		// Delete consumer
		//auctionService.DeleteConsumer(auctionID, c)
		return
	}

	log.Printf("\nSTEP: 14: RPEPARE AUCTION MEMBERS\n\n")
	response := responses.AuctionMembers(members, userAndPrice)
	if err = c.WriteJSON(response); err != nil && err == c.Close() {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send data with error: %s", err)))

		// Delete consumer
		auctionService.DeleteConsumer(auctionID, c)
		return
	}

	log.Printf("\nSTEP: 15: SET ACTUAL IN STACK\n\n")
	auctionService.SetActual(auctionID, c)

	//UpdateOwner(c, ctx, auctionID, userID)
	return
}
