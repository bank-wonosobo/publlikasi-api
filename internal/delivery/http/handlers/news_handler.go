package handlers

import (
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/request"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/response"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type NewsHandler struct {
	newsService services.NewsService
	validator   validator.Validator
}

func NewNews(newsService services.NewsService, validator validator.Validator) *NewsHandler {
	return &NewsHandler{
		newsService: newsService,
		validator:   validator,
	}
}

func (h *NewsHandler) Create(c *fiber.Ctx) error {
	var request request.NewsCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.newsService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}
