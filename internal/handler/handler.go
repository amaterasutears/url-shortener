package handler

import (
	"github.com/amaterasutears/url-shortener/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LinkService interface {
	Shorten(original string) (string, error)
	Redirect(code string) (string, error)
}

func Shorten(linkService LinkService, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p dto.ShortenQueryParam

		err := c.QueryParser(&p)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err = validate.Struct(&p)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		code, err := linkService.Shorten(p.Original)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).SendString(code)
	}
}

func Redirect(linkService LinkService, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p dto.RedirectParam

		err := c.ParamsParser(&p)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err = validate.Struct(&p)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		original, err := linkService.Redirect(p.Code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Redirect(original, fiber.StatusFound)
	}
}