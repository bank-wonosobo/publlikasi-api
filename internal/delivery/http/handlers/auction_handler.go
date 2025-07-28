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
