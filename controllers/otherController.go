package controllers

import "github.com/gofiber/fiber/v2"

func Other(c *fiber.Ctx) error {
	return c.SendString("Hello, Other World!")
}
