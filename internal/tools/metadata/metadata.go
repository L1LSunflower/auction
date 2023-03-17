package metadata

import (
	"fmt"
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
	if metadata.PerPage, err = ParseParamInt(args, PerPageParams); err != nil {
		return nil, errorhandler.ErrGettingPerPage
	}

	if metadata.CurrentPage, err = ParseParamInt(args, PageParams); err != nil {
		return nil, errorhandler.ErrGettingPage
	}

	if metadata.PerPage > perPageConst {
		return nil, errorhandler.ErrWrongPerPageValue
	}

	return metadata, nil
}

func PrepareTags(tags string) string {
	tags = strings.ReplaceAll(tags, "_", " ")
	tagsSlice := strings.Split(tags, ",")
	tags = ""
	for iTag, tag := range tagsSlice {
		if iTag == len(tagsSlice)-1 {
			tags += fmt.Sprintf("'%s'", tag)
		} else {
			tags += fmt.Sprintf("'%s',", tag)
		}
	}

	return tags
}

func ParseParamInt(args *fasthttp.Args, nameParams string) (int, error) {
	v, err := strconv.Atoi(string(args.Peek(nameParams)))
	if err != nil {
		return 0, err
	}
	return v, nil

}

func ParseParams(args *fasthttp.Args, nameParams string) string {
	return string(args.Peek(nameParams))
}

func ConcatStrings(sliceOfStrings []string, delimiter string) string {
	return strings.Join(sliceOfStrings, delimiter)
}
