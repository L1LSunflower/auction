package cities

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type CityInterface interface {
	Create(ctx context.Context, name string) (*entities.City, error)
	Get(ctx context.Context, id int) (*entities.City, error)
	GetAll(ctx context.Context) ([]*entities.City, error)
	Update(ctx context.Context, id int, name string) (int, error)
	Delete(ctx context.Context, id int) error
}

func NewRepository() CityInterface {
	return &Repository{}
}
