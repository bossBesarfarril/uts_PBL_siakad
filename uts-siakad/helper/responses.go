package helper

import (
	"uts-siakad/app/model"

	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(model.WebResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessList(c *fiber.Ctx, message string, data interface{}, meta *model.Meta) error {
	return c.Status(200).JSON(model.WebResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(204)
}
