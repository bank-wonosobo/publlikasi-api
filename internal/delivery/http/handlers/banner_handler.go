package handlers

import (
	"strings"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type BannerHandler struct {
	bannerService services.BannerService
	validator     validator.Validator
}

func NewBanner(bannerService services.BannerService,
	validator validator.Validator) *BannerHandler {
	return &BannerHandler{
		bannerService: bannerService,
		validator:     validator,
	}
}

func (h *BannerHandler) Create(c *fiber.Ctx) error {
	var request dto.BannerCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate file
	err = h.validator.ValidateImageFile(file)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.bannerService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *BannerHandler) Index(c *fiber.Ctx) error {
	// parse params
	var params dto.BannerGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get status publish
	if strings.Split(c.Path(), "/")[3] != "admin" {
		isActive := true
		params.IsActive = &isActive
	}
	// call service
	result, total, err := h.bannerService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	// calculate total page
	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	// return result
	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}
func (h *BannerHandler) Update(c *fiber.Ctx) error {
	var request dto.BannerUpdateRequest
	id := c.Params("id")
	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("image")
	if request.Image != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
		}
	}

	// call service
	result, err := h.bannerService.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *BannerHandler) Delete(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	err := h.bannerService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}

func (h *BannerHandler) Activate(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.bannerService.Activate(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}
