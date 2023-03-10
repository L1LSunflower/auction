package responses

import (
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	tagsResponse "github.com/L1LSunflower/auction/internal/responses/structs"
)

func Tags(ctx *fiber.Ctx, tags []*entities.Tag) error {
	var tagsInString []string
	for _, tag := range tags {
		tagsInString = append(tagsInString, tag.Name)
	}

	return ctx.JSON(tagsResponse.Tags{
		Status: successStatus,
		Tags:   tagsInString,
	})
}
