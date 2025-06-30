package handlers

import (
	"strconv"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ComplaintTypeHandler struct {
	complaintTypeService services.ComplaintTypeService
	validator            validator.Validator
}

func NewComplaintType(
	complaintTypeService services.ComplaintTypeService,
	validator validator.Validator) *ComplaintTypeHandler {
	return &ComplaintTypeHandler{
		complaintTypeService: complaintTypeService,
		validator:            validator,
	}
}

func (h *ComplaintTypeHandler) Index(c *fiber.Ctx) error {
	// call service
	result, err := h.complaintTypeService.Index(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

// create handler
func (h *ComplaintTypeHandler) Create(c *fiber.Ctx) error {
	var request dto.ComplaintTypeCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.complaintTypeService.Create(c.Context(), &request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ComplaintTypeHandler) Update(c *fiber.Ctx) error {
	// get id
	idParams := c.Params("id")
	// Convert string to int
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid complaint type ID")
	}

	var request dto.ComplaintTypeCreateRequest
	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.complaintTypeService.Update(c.Context(), &request, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ComplaintTypeHandler) Delete(c *fiber.Ctx) error {
	// get id
	idParams := c.Params("id")
	// Convert string to int
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid complaint type ID")
	}

	err = h.complaintTypeService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}
