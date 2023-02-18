package handlers

import "github.com/gofiber/fiber/v2"

func Healthcheck(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}
