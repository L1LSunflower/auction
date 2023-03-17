package auction_websockets

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strconv"

	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	auctionReq "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/redisdb"
)

func Auction(c *websocket.Conn) {
	//defer c.Close()

	for {

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
		redisConn := redisdb.RedisInstance().RedisClient

		ctx, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
		if err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		auction, err := db_repository.AuctionInterface.Auction(ctx, id)
		if err != nil || auction == nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		if auction.Status != entities.ActiveStatus {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		auctionOffer := &auctionReq.AmountOffer{}
		if err = c.ReadJSON(auctionOffer); err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to parser request"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		member, err := db_repository.AuctionInterface.Member(ctx, auction.ID, auctionOffer.ID)
		if err != nil || member == nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to get auction"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		if auctionOffer.Amount < auction.Price {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "your amount lower than that"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		if err = db_repository.AuctionInterface.UpdatePrice(ctx, auction.ID, auctionOffer.Amount); err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to set price for some reason"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}

		if err = c.WriteJSON(fiber.Map{"status": "success", "price": auctionOffer.Amount}); err != nil {
			if err = c.WriteJSON(fiber.Map{"status": "error", "message": "failed to set price for some reason"}); err != nil {
				log.Println("ERROR: failed to get auction with that id")
				return
				//break
			}
		}
	}
}
