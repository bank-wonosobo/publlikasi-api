package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/request"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/response"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	reportService services.ReportService
	validator     validator.Validator
}

func NewReport(reportService services.ReportService, validator validator.Validator) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
		validator:     validator,
	}
}

func (h *ReportHandler) Create(c *fiber.Ctx) error {
	var request request.ReportCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.reportService.Create(c.Context(), &request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}
