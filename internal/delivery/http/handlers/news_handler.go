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

func (h *NewsHandler) GetAll(c *fiber.Ctx) error {
	// parse params
	var params request.NewsGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
	}

	// call service
	result, total, err := h.newsService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	// calculate total page
	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	// return result
	return c.Status(fiber.StatusOK).JSON(response.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *NewsHandler) Update(c *fiber.Ctx) error {
	var request request.NewsUpdateRequest
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
	file, err := c.FormFile("image")
	if request.Image != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateReponseError(err.Error()))
		}
	}

	// call service
	result, err := h.newsService.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *NewsHandler) Delete(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	err := h.newsService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(id))
}

func (h *NewsHandler) Approve(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.newsService.Approve(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *NewsHandler) Archive(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.newsService.Archive(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *NewsHandler) Detail(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.newsService.Detail(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}

func (h *NewsHandler) DetailWSlug(c *fiber.Ctx) error {
	// get id
	id := c.Params("slug")

	// call service
	result, err := h.newsService.DetailBySlug(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateReponseSuccess(result))
}
