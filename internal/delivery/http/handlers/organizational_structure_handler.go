package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type OrganizationalStructureHandler struct {
	service   services.OrganizationalStructureService
	validator validator.Validator
}

func NewOrganizationalStructure(service services.OrganizationalStructureService, validator validator.Validator) *OrganizationalStructureHandler {
	return &OrganizationalStructureHandler{
		service:   service,
		validator: validator,
	}
}

func (h *OrganizationalStructureHandler) Create(c *fiber.Ctx) error {
	var request dto.OrganizationalStructureCreateReq

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

	result, err := h.service.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *OrganizationalStructureHandler) GetAll(c *fiber.Ctx) error {
	var params dto.OrganizationalStructureGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	result, total, err := h.service.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *OrganizationalStructureHandler) Update(c *fiber.Ctx) error {
	var request dto.OrganizationalStructureUpdateReq
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

	result, err := h.service.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *OrganizationalStructureHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}
