package auctions

import (
	"context"
	"fmt"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
)

func EndAuctions(ctx context.Context, timeInterval time.Duration) {
	for {
		if err := db_repository.AuctionInterface.EndAuctions(ctx); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to change status for auctions from %s to %s, with error: %s", entities.ActiveStatus, entities.CompletedStatus, err.Error())))
		}
		time.Sleep(timeInterval)
	}
}
