package handlers

import (
	"strconv"
	"strings"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/services"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	reportService services.ReportService
	validator     validator.Validator
}

func NewReport(reportService services.ReportService, validator validator.Validator) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
		validator:     validator,
	}
}

func (h *ReportHandler) Index(c *fiber.Ctx) error {
	// parse params
	var params dto.ReportGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get status publish
	if strings.Split(c.Path(), "/")[3] != "admin" {
		params.Status = "published"
	}

	// call service
	result, total, err := h.reportService.Index(c.Context(), &params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	// calculate total page
	totalPage := (total + int64(params.Limit) - 1) / int64(params.Limit)

	// return result
	return c.Status(fiber.StatusOK).JSON(dto.CreatePaginateResponse(result, params.Page, params.Limit, total, totalPage))
}

func (h *ReportHandler) Create(c *fiber.Ctx) error {
	var request dto.ReportCreateRequest

	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// get file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}
	// call service
	result, err := h.reportService.Create(c.Context(), &request, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportHandler) UploadFile(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// get file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.reportService.UploadFile(c.Context(), file, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportHandler) Update(c *fiber.Ctx) error {
	var request dto.ReportUpdateRequest
	id := c.Params("id")
	// parse request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// validate request
	if err := h.validator.Validate(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.reportService.Update(c.Context(), &request, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportHandler) Delete(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	err := h.reportService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(id))
}

func (h *ReportHandler) Approve(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.reportService.Approve(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportHandler) Archive(c *fiber.Ctx) error {
	// get id
	id := c.Params("id")

	// call service
	result, err := h.reportService.Archive(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}

func (h *ReportHandler) GetByReportType(c *fiber.Ctx) error {
	// get report type id
	idParams := c.Params("report_type_id")

	reportTypeID, err := strconv.Atoi(idParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}
	// parse params
	var params dto.ReportGetQueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateReponseError(err.Error()))
	}

	// call service
	result, err := h.reportService.GetByReportType(c.Context(), &params, reportTypeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateReponseError(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateReponseSuccess(result))
}
