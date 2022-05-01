package api

import (
	"github.com/gofiber/fiber/v2"
	"merchant-account-api/core/domain"
	"merchant-account-api/ports/api"
	"merchant-account-api/utils"
	"strconv"
	"time"
)

type Response struct {
	Code        uint
	Description string
}

type MerchantAdapter struct {
	service api.MerchantServiceInterface
}

func NewMerchantAdapter(service api.MerchantServiceInterface) api.HttpHandler {
	return &MerchantAdapter{service: service}
}

// Post func Updates a merchant
// @Description Updates a merchant
// @Summary Updates a merchant
// @Tags Merchant
// @Accept json
// @Produce json
// @Param merchant_code body string true "Merchant code"
// @Param merchant_name body string true "Merchant name"
// @Param merchant_email body string true "Merchant email"
// @Param merchant_city body string true "Merchant city"
// @Param merchant_category_code body string true "Merchant category code"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/merchant [post]
func (m MerchantAdapter) Post(ctx *fiber.Ctx) error {
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
		MerchantCode         string `json:"merchant_code"`
		MerchantName         string `json:"merchant_name"`
		MerchantEmail        string `json:"merchant_email"`
		MerchantCity         string `json:"merchant_city"`
		MerchantCategoryCode string `json:"merchant_category_code"`
	}{}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating merchant",
		})
	}
	merchantCodeInt, err := strconv.Atoi(payload.MerchantCode)
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid updating merchant",
		})
	}
	err = m.service.Update(domain.Merchant{
		MerchantCode:         uint32(merchantCodeInt),
		MerchantName:         payload.MerchantName,
		MerchantEmail:        payload.MerchantEmail,
		MerchantCity:         payload.MerchantCity,
		MerchantCategoryCode: payload.MerchantCategoryCode,
	})
	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid updating merchant",
		})
	}
	return ctx.JSON(Response{
		Code:        200,
		Description: "Successfully updated a merchant",
	})
}

// Put func Creates a merchant
// @Description Create a merchant
// @Summary Creates a merchant
// @Tags Merchant
// @Accept json
// @Produce json
// @Param merchant_name body string true "Merchant name"
// @Param merchant_email body string true "Merchant email"
// @Param merchant_city body string true "Merchant city"
// @Param merchant_category_code body string true "Merchant category code"
// @Success 201
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/merchant/ [put]
func (m MerchantAdapter) Put(ctx *fiber.Ctx) error {
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
		MerchantName         string `json:"merchant_name"`
		MerchantEmail        string `json:"merchant_email"`
		MerchantCity         string `json:"merchant_city"`
		MerchantCategoryCode string `json:"merchant_category_code"`
	}{}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating merchant",
		})
	}
	merchantCode, err := m.service.Save(domain.Merchant{
		MerchantName:         payload.MerchantName,
		MerchantEmail:        payload.MerchantEmail,
		MerchantCity:         payload.MerchantCity,
		MerchantCategoryCode: payload.MerchantCategoryCode,
	})
	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid creating merchant",
		})
	}
	return ctx.JSON(Response{
		Code:        201,
		Description: "Successfully created a merchant with id: " + merchantCode,
	})
}

// Delete func Removes a merchant.
// @Description Removes a merchant.
// @Summary Removes a merchant
// @Tags Merchant
// @Accept json
// @Produce json
// @Param merchant_code path string true "Merchant Code"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/merchant/{merchant_code} [delete]
func (m MerchantAdapter) Delete(ctx *fiber.Ctx) error {
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
	code := ctx.Params("merchant_code")
	err = m.service.Delete(code)

	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid merchant ID",
		})
	}
	return ctx.JSON(Response{
		Code:        200,
		Description: "Successfully delete a merchant",
	})
}

// Get func Gets details of a merchant
// @Description Gets details of a merchant (name, city, email, mcc)
// @Summary Gets details of a merchant
// @Tags Merchant
// @Accept json
// @Produce json
// @Param merchant_code path string true "Merchant Code"
// @Success 200 {object} domain.Merchant
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/merchant/{merchant_code} [get]
func (m MerchantAdapter) Get(ctx *fiber.Ctx) error {
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
	code := ctx.Params("merchant_code")
	result, err := m.service.Find(code)

	if err != nil {
		return ctx.JSON(Response{
			Code:        400,
			Description: "Invalid merchant ID",
		})
	}
	return ctx.JSON(result)
}
