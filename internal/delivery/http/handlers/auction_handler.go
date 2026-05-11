package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type AuctionHandler struct {
	auctionService services.AuctionService
	validator      validator.Validator
}

func NewAuction(auctionService services.AuctionService,
	validator validator.Validator) *AuctionHandler {
	return &AuctionHandler{
		auctionService: auctionService,
		validator:      validator,
	}
}

func (h *AuctionHandler) Create(c *fiber.Ctx) error {
	var request dto.AuctionCreateRequest

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
	result, err := h.auctionService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *AuctionHandler) Index(c *fiber.Ctx) error {
	var params dto.AuctionGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	result, total, err := h.auctionService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)
	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *AuctionHandler) Update(c *fiber.Ctx) error {
	var request dto.AuctionUpdateRequest
	id := c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	file, err := c.FormFile("image")
	if request.Image != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
		}
		err = h.validator.ValidateImageFile(file)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
		}
	}

	result, err := h.auctionService.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *AuctionHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.auctionService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}
