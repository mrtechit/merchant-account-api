package api

import (
	"github.com/gofiber/fiber/v2"
	"merchant-account-api/core/domain"
	"merchant-account-api/ports/api"
	"merchant-account-api/utils"
	"strconv"
	"time"
)

type TeamMemberResponse struct {
	id    uint32
	Name  string
	Email string
}

type TeamMemberAdapter struct {
	service api.TeamMembersServiceInterface
}

func NewTeamMemberAdapter(service api.TeamMembersServiceInterface) api.TeamMembersHandler {
	return &TeamMemberAdapter{service: service}
}

// Get func Gets details of a team member
// @Description Gets details of a team member
// @Summary Gets details of a team member
// @Tags Team Members
// @Accept json
// @Produce json
// @Param id path string true "team member"
// @Success 200 {object} domain.TeamMember
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/teams/{id} [get]
func (t *TeamMemberAdapter) Get(ctx *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires := claims.Expires
	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	id := ctx.Params("id")
	result, err := t.service.Find(id)

	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid team member ID",
		})
	}
	response := TeamMemberResponse{
		id:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}
	return ctx.JSON(response)
}

// Post func update team member
// @Description Updates a team member
// @Summary updates a team member
// @Tags Team members
// @Accept json
// @Produce json
// @Param id body string true "Team member id"
// @Param name body string true "Team member name"
// @Param email body string true "Team member email"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/teams [post]
func (t *TeamMemberAdapter) Post(ctx *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires := claims.Expires
	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	payload := struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid updating team member",
		})
	}

	Id, err := strconv.Atoi(payload.Id)
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating team member",
		})
	}

	err = t.service.Update(domain.TeamMember{
		ID:    uint32(Id),
		Name:  payload.Name,
		Email: payload.Email,
	})
	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid updating team member",
		})
	}
	return ctx.JSON(Response{
		Code:        200,
		Description: "Successfully updated a team member",
	})
}

// Put func creates a team member
// @Description Creates a team member
// @Summary creates a team member
// @Tags Team members
// @Accept json
// @Produce json
// @Param name body string true "Team member name"
// @Param email body string true "Team member email"
// @Param merchant_code body string true "Merchant code for this member"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/teams/ [put]
func (t *TeamMemberAdapter) Put(ctx *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires := claims.Expires
	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	payload := struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		MerchantMerchantCode string `json:"merchant_code"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating team member",
		})
	}
	merchantCodeInt, err := strconv.Atoi(payload.MerchantMerchantCode)
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating team member",
		})
	}
	err = t.service.Save(domain.TeamMember{
		Name:                 payload.Name,
		Email:                payload.Email,
		MerchantMerchantCode: uint32(merchantCodeInt),
	})
	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating team member",
		})
	}
	return ctx.JSON(Response{
		Code:        201,
		Description: "Successfully created a team member",
	})
}

// Delete func removes a team member.
// @Description Remove a team member.
// @Summary remove a team member
// @Tags Team members
// @Accept json
// @Produce json
// @Param id path string true "Team member id"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/teams/{id} [delete]
func (t *TeamMemberAdapter) Delete(ctx *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires := claims.Expires
	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	id := ctx.Params("id")
	err = t.service.Delete(id)
	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid merchant team member",
		})
	}
	return ctx.JSON(Response{
		Code:        200,
		Description: "Successfully delete a team member",
	})
}

// GetTeamMembersByMerchantCode func gets all team members for a particular merchant
// @Description Gets all team members for a particular merchant
// @Summary gets all team members for a particular merchant
// @Tags Team members
// @Accept json
// @Produce json
// @Param merchant_code path string true "Merchant code"
// @Param page path string true "Page"
// @Param page_size path string true "Page Size"
// @Success 200 {array} domain.TeamMember
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/teams/{merchant_code}/{page}/{page_size} [get]
func (t *TeamMemberAdapter) GetTeamMembersByMerchantCode(ctx *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires := claims.Expires
	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	merchantCode := ctx.Params("merchant_code")
	page, _ := strconv.Atoi(ctx.Params("page"))
	pageSize, _ := strconv.Atoi(ctx.Params("page_size"))
	teamMembers, err := t.service.GetTeamMembersByMerchantCode(merchantCode, page, pageSize)
	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid get team members by merchant",
		})
	}
	if len(teamMembers) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Team members not found for merchant",
		})
	}
	return ctx.JSON(teamMembers)
}
