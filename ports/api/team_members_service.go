package api

import "github.com/gofiber/fiber/v2"

type TeamMembersHandler interface {
	HttpHandler
	GetTeamMembersByMerchantCode(*fiber.Ctx) error
}
