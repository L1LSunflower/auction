package tags

import (
	"context"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	tagsRequest "github.com/L1LSunflower/auction/internal/requests/structs/tags"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func ByPattern(ctx context.Context, tag *tagsRequest.Tag) ([]*entities.Tag, error) {
	tags, err := db_repository.TagsInterface.TagsByPattern(ctx, tag.Pattern)
	if err != nil {
		return nil, errorhandler.ErrGetTags
	}

	if len(tags) <= 0 {
		return []*entities.Tag{}, nil
	}

	return tags, nil
}
