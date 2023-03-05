package files

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type FilesInterface interface {
	Create(ctx context.Context, filename string) (*entities.File, error)
	CreateLink(ctx context.Context, itemID int, filename string) (*entities.ItemFiles, error)
	Files(ctx context.Context, itemID int) ([]*entities.File, error)
	Delete(ctx context.Context, itemID int, filename string) error
	DeleteAll(ctx context.Context, itemId int) error
}

func GetFilesInterface() FilesInterface {
	return &Repository{}
}
