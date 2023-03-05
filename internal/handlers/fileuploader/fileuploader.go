package fileuploader

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func UploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	uniqueId, err := uuid.NewV4()
	if err != nil {
		return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	name := strings.Replace(uniqueId.String(), "-", "", -1)
	splFile := strings.Split(file.Filename, ".")
	fileExt := splFile[len(splFile)-1]
	filename := fmt.Sprintf("%s.%s", name, fileExt)

	if err = ctx.SaveFile(file, fmt.Sprintf("./static/%s", filename)); err != nil {
		return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	data := map[string]interface{}{
		"content_type": file.Header.Get("Content-Type"),
		"file_name":    filename,
	}

	return ctx.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func DeleteFile(c *fiber.Ctx) error {
	fileName := c.Params("fileName")

	if err := os.Remove(fmt.Sprintf("./images/%s", fileName)); err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
