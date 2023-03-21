package auctions

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/pkg/logger/message"
	"time"

	"github.com/L1LSunflower/auction/pkg/logger"
)

func StartAuctions(ctx context.Context, timeInterval time.Duration) {
	for {
		if err := db_repository.AuctionInterface.ActivateAuctions(ctx); err != nil {
			logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to change status for auctions from %s to %s, with error: %s", entities.InactiveStatus, entities.ActiveStatus, err.Error())))
		}
		time.Sleep(timeInterval)
	}
}
