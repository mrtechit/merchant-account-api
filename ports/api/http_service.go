package api

import "github.com/gofiber/fiber/v2"

type HttpHandler interface {
	Get(*fiber.Ctx) error
	Post(*fiber.Ctx) error
	Put(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}
