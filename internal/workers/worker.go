package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	auctionsWorkers "github.com/L1LSunflower/auction/internal/workers/auctions"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
	"github.com/L1LSunflower/auction/pkg/redisdb"
)

const (
	StartAuctionsInterval = 3 * time.Minute
	EndAuctionsInterval   = 1 * time.Minute
)

func StartWorkers() {
	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisClient := redisdb.RedisInstance().RedisClient

	ctx, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisClient)
	if err != nil {
		logger.Log.Fatal(message.NewMessage(fmt.Sprintf("failed to start context with depends with error: %s", err.Error())))
	}

	go auctionsWorkers.StartAuctions(ctx, StartAuctionsInterval)
	go auctionsWorkers.EndAuctions(ctx, EndAuctionsInterval)
}
