package handlers

import (
	"strconv"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ReportTypeHandler struct {
	reportTypeService services.ReportTypeService
	validator         validator.Validator
}

func NewReportType(reportTypeService services.ReportTypeService, validator validator.Validator) *ReportTypeHandler {
	return &ReportTypeHandler{
		reportTypeService: reportTypeService,
		validator:         validator,
	}
}

func (h *ReportTypeHandler) Index(c *fiber.Ctx) error {
	// call service
	result, err := h.reportTypeService.Index(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

// create handler
func (h *ReportTypeHandler) Create(c *fiber.Ctx) error {
	var request dto.ReportTypeCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.reportTypeService.Create(c.Context(), &request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportTypeHandler) Update(c *fiber.Ctx) error {
	// get id
	idParams := c.Params("id")
	// Convert string to int
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	var request dto.ReportTypeCreateRequest
	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.reportTypeService.Update(c.Context(), &request, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportTypeHandler) Delete(c *fiber.Ctx) error {
	// get id
	idParams := c.Params("id")
	// Convert string to int
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	err = h.reportTypeService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}
