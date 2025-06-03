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

func (h *AnnouncementHandler) Create(c *fiber.Ctx) error {
	var request dto.AnnouncementCreateReq

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("attachment")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.service.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *AnnouncementHandler) Update(c *fiber.Ctx) error {
	var request dto.AnnouncementUpdateReq
	id := c.Params("id")
	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("attachment")
	if request.Attachment != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
		}
	}

	// call service
	result, err := h.service.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *AnnouncementHandler) Delete(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	err := h.service.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(id))
}

func (h *AnnouncementHandler) Approve(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.service.Approve(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *AnnouncementHandler) Archive(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.service.Archive(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *AnnouncementHandler) Detail(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.service.Detail(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}
