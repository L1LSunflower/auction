package metadata

import (
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const (
	perPageConst = 30
)

var (
	PerPageParams = "per_page"
	PageParams    = "page"
)

type Metadata struct {
	CurrentPage int
	LastPage    int
	PerPage     int
	Total       int
	Limit       int
	Offset      int
}

func GetParams(ctx *fiber.Ctx) (*Metadata, error) {
	var metadata = &Metadata{}
	var err error
	args := ctx.Request().URI().QueryArgs()

	// Metadata
	if metadata.PerPage, err = parseParam(args, PerPageParams); err != nil {
		return nil, errorhandler.ErrGettingPerPage
	}

	if metadata.CurrentPage, err = parseParam(args, PageParams); err != nil {
		return nil, errorhandler.ErrGettingPage
	}

	if metadata.PerPage > perPageConst {
		return nil, errorhandler.ErrWrongPerPageValue
	}

	return metadata, nil
}

func parseParam(args *fasthttp.Args, nameParams string) (int, error) {
	v, err := strconv.Atoi(string(args.Peek(nameParams)))
	if err != nil {
		return 0, err
	}
	return v, nil

}

func ConcatStrings(sliceOfStrings []string, delimiter string) string {
	return strings.Join(sliceOfStrings, delimiter)
}
