package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ComplaintHandler struct {
	complaintService services.ComplaintService
	validator        validator.Validator
}

func NewComplaint(complaintService services.ComplaintService,
	validadator validator.Validator) *ComplaintHandler {
	return &ComplaintHandler{
		complaintService: complaintService,
		validator:        validadator,
	}
}

func (h *ComplaintHandler) Index(c *fiber.Ctx) error {
	// parse params
	var params dto.ComplaintGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, total, err := h.complaintService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	// calculate total page
	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	// return result
	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *ComplaintHandler) Create(c *fiber.Ctx) error {
	var request dto.ComplaintCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("evidence_file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate file
	err = h.validator.ValidateImageFile(file)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.complaintService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ComplaintHandler) Process(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.complaintService.Process(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ComplaintHandler) Done(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.complaintService.Done(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}
