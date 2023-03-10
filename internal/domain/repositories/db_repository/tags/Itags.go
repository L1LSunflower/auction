package tags

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type TagsInterface interface {
	Create(ctx context.Context, tagName string) (*entities.Tag, error)
	ByName(ctx context.Context, tagName string) (*entities.Tag, error)
	CreateLink(ctx context.Context, itemID, fileID int) (*entities.ItemTags, error)
	Tags(ctx context.Context, itemID int) ([]*entities.Tag, error)
	DeleteItemTags(ctx context.Context, itemID, tagID int) error
	DeleteItemLinks(ctx context.Context, itemID int) error
	TagsByPattern(ctx context.Context, pattern string) ([]*entities.Tag, error)
}

func GetTagsInterface() TagsInterface {
	return &Repository{}
}
