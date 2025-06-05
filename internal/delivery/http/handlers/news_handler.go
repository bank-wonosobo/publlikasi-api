package handlers

import (
	"strings"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
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
	var request dto.NewsCreateRequest

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
	result, err := h.newsService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *NewsHandler) GetAll(c *fiber.Ctx) error {
	// parse params
	var params dto.NewsGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get status publish
	if strings.Split(c.Path(), "/")[3] != "admin" {
		params.Status = "published"
	}
	// call service
	result, total, err := h.newsService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	// calculate total page
	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	// return result
	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *NewsHandler) Update(c *fiber.Ctx) error {
	var request dto.NewsUpdateRequest
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
		// validate file
		err = h.validator.ValidateImageFile(file)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
		}
	}

	// call service
	result, err := h.newsService.Update(c.Context(), &request, file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *NewsHandler) Delete(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	err := h.newsService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}

func (h *NewsHandler) Approve(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.newsService.Approve(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *NewsHandler) Archive(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.newsService.Archive(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *NewsHandler) Detail(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.newsService.Detail(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *NewsHandler) DetailWSlug(c *fiber.Ctx) error {
	// get id
	id := c.Params("slug")

	// call service
	result, err := h.newsService.DetailBySlug(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}
