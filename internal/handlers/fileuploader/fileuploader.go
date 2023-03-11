package fileuploader

import (
	"fmt"
	"github.com/kolesa-team/go-webp/webp"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

const (
	fileExtension = "webp"
)

func UploadFile(ctx *fiber.Ctx) error {
	fileRequest, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(fiber.Map{"status": "error", "message": "Server error", "data": nil})

	}

	uniqueId, err := uuid.NewV4()
	if err != nil {
		return ctx.JSON(fiber.Map{"status": "error", "message": "Server error", "data": nil})
	}

	name := strings.Replace(uniqueId.String(), "-", "", -1)
	splFile := strings.Split(fileRequest.Filename, ".")
	fileExt := splFile[len(splFile)-1]
	filename := fmt.Sprintf("%s.%s", name, fileExt)

	if fileExt == "mp4" {
		if err = ctx.SaveFile(fileRequest, fmt.Sprintf("%s/%s", "static", filename)); err != nil {
			return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}

		data := map[string]interface{}{
			"content_type": fileRequest.Header.Get("Content-Type"),
			"file_name":    filename,
		}

		return ctx.JSON(fiber.Map{"status": 201, "message": "video uploaded successfully", "data": data})
	}

	filename = fmt.Sprintf("%s.%s", name, fileExtension)
	file, err := fileRequest.Open()
	if err != nil {
		return ctx.JSON(fiber.Map{"status": "error", "message": "Server error", "data": nil})
	}

	var img image.Image
	switch fileExt {
	case "jpeg":
		img, err = jpeg.Decode(file)
	case "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		return ctx.JSON(fiber.Map{"status": "error", "message": "non-extendable file type", "data": nil})
	}
	if err != nil {
		return ctx.JSON(fiber.Map{"status": "error", "message": "Server error", "data": nil})
	}

	output, err := os.Create(fmt.Sprintf("%s/%s", "static", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	if err = webp.Encode(output, img, nil); err != nil {
		log.Fatalln(err)
	}

	data := map[string]interface{}{
		"content_type": "image/webp",
		"file_name":    filename,
	}

	return ctx.JSON(fiber.Map{"status": 201, "message": "image uploaded successfully", "data": data})
}

func DeleteFile(c *fiber.Ctx) error {
	fileName := c.Params("file")

	if err := os.Remove(fmt.Sprintf("%s/%s", "static", fileName)); err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
