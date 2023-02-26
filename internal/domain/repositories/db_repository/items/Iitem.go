package items

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type ItemInterface interface {
	Create(ctx context.Context, item *entities.Item) error
	Item(ctx context.Context, id int) (*entities.Item, error)
	Update(ctx context.Context, item *entities.Item) error
	Delete(ctx context.Context, id int) error
}

func GetItemInterface() ItemInterface {
	return &Repository{}
}
