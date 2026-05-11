package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	profileService services.ProfileService
	validator      validator.Validator
}

func NewProfile(profileService services.ProfileService, validator validator.Validator) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		validator:      validator,
	}
}

func (h *ProfileHandler) Create(c *fiber.Ctx) error {
	var request dto.ProfileCreateReq

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	file, err := c.FormFile("image")
	if err != nil {
		file = nil
	}

	if file != nil {
		err = h.validator.ValidateImageFile(file)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
		}
	}

	result, err := h.profileService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ProfileHandler) GetAll(c *fiber.Ctx) error {
	var params dto.ProfileGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	result, total, err := h.profileService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *ProfileHandler) Update(c *fiber.Ctx) error {
	var request dto.ProfileUpdateReq
	id := c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	file, err := c.FormFile("image")
	if err != nil {
		file = nil
	}

	if file != nil {
		err = h.validator.ValidateImageFile(file)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
		}
	}

	result, err := h.profileService.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ProfileHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.profileService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}
