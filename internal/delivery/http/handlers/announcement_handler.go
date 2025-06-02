package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/response"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type AnnouncementHandler struct {
	service   services.AnnouncementService
	validator validator.Validator
}

func NewAnnouncement(service services.AnnouncementService,
	validator validator.Validator) *AnnouncementHandler {
	return &AnnouncementHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AnnouncementHandler) Index(c *fiber.Ctx) error {
	// parse params
	var params dto.AnnouncementGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// call service
	result, total, err := h.service.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	// calculate total page
	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	// return result
	return c.Status(fiber.StatusOK).JSON(response.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}
