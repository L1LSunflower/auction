package responses

import (
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/gofiber/fiber/v2"
	"time"

	balanceResp "github.com/L1LSunflower/auction/internal/responses/structs"
)

func Credit(ctx *fiber.Ctx, balance *aggregates.UserBalance) error {
	return ctx.JSON(&balanceResp.Credit{
		Status:  successStatus,
		ID:      balance.Balance.ID,
		Balance: balance.Balance.Balance,
		Date:    time.Now().Format(dateFormat),
	})
}

func Debit(ctx *fiber.Ctx, balance *aggregates.UserBalance) error {
	return ctx.JSON(&balanceResp.Credit{
		Status:  successStatus,
		ID:      balance.Balance.ID,
		Balance: balance.Balance.Balance - 1000,
		Date:    time.Now().Format(dateFormat),
	})
}

func Balance(ctx *fiber.Ctx, balance *aggregates.UserBalance) error {
	return ctx.JSON(&balanceResp.Credit{
		Status:  successStatus,
		ID:      balance.Balance.ID,
		Balance: balance.Balance.Balance,
		Date:    time.Now().Format(dateFormat),
	})
}
