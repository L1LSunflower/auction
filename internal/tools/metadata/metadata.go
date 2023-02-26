package metadata

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const perPageConst = 30

type Metadata struct {
	CurrentPage int
	LastPage    int
	PerPage     int
	Total       int
	Limit       int
	Offset      int
}

func GetParams(ctx *fiber.Ctx) (*Metadata, error) {
	var metadata Metadata
	var err error
	args := ctx.Request().URI().QueryArgs()

	metadata.CurrentPage, err = parseParam(args, "page")
	if err != nil {
		return nil, fmt.Errorf("failed to get params 'page' with error: %s", err.Error())
	}

	metadata.PerPage, err = parseParam(args, "per_page")
	if err != nil {
		return nil, fmt.Errorf("failed to get params 'per_page' with error: %s", err.Error())
	}

	if metadata.PerPage > perPageConst {
		return nil, fmt.Errorf("per_page params must be less than %d", perPageConst)
	}

	return &metadata, nil
}

func parseParam(args *fasthttp.Args, nameParams string) (int, error) {
	v, err := strconv.Atoi(string(args.Peek(nameParams)))
	if err != nil {
		return 0, err
	}
	return v, nil

}

func Filter(ctx *fiber.Ctx) ([]string, error) {
	var where = []string{}

	for i := 1; i <= 10; i++ {
		if tag := ctx.Get(fmt.Sprintf("tag%d", i)); tag != "" {
			where = append(where, fmt.Sprintf("tag%d", i)+"="+tag)
		}
	}

	if isActive := ctx.Get("is_active"); isActive == "true" {
		where = append(where, "is_active=1")
	}

	if category := ctx.Get("category"); category != "" {
		where = append(where, fmt.Sprintf("category=%s", category))
	}

	return where, nil
}
