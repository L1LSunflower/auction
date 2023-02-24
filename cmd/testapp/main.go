package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	// create new fiber instance  and use across whole app
	app := fiber.New(fiber.Config{
		//Prefork:                 true,
		BodyLimit:               100 * 1024 * 1024,
		ServerHeader:            "accept",
		StrictRouting:           true,
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"http://localhost:3000", "http://localhost:8100", "http://192.168.0.15:3000", "http://192.168.0.14:8100"},
		RequestMethods:          []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
	})

	// middleware to allow all clients to communicate using http and allow cors
	app.Use(cors.New())

	// serve  images from images directory prefixed with /images
	// i.e http://localhost:4000/images/someimage.webp

	app.Static("/images", "./images")

	// handle image uploading using post request

	app.Post("/", handleFileupload)

	// delete uploaded image by providing unique image name

	app.Delete("/:imageName", handleDeleteImage)

	// start dev server on port 3000

	log.Fatal(app.Listen("0.0.0.0:3000"))
}

func handleFileupload(c *fiber.Ctx) error {

	// parse incomming image file

	file, err := c.FormFile("file")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("images/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func handleDeleteImage(c *fiber.Ctx) error {
	// extract image name from params
	imageName := c.Params("imageName")

	// delete image from ./images
	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
