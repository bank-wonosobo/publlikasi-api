package dto

import "time"

// request
type ReportCreateRequest struct {
	Title       string    `json:"title" validate:"required" form:"title"`
	Description string    `json:"description" validate:"required" form:"description"`
	PeriodStart time.Time `json:"period_start" validate:"required" form:"period_start"`
	PeriodEnd   time.Time `json:"period_end" validate:"required" form:"period_end"`
	Year        int       `json:"year" validate:"required" form:"year"`
	Quarter     *int      `json:"quarter" form:"quarter"`
	Version     string    `json:"version" validate:"required" form:"version"`
	ReportType  string    `json:"report_type" validate:"required" form:"report_type"`
	File        string    `json:"file" form:"file"`
}

type ReportGetQueryParams struct {
	Title       string `query:"title"`
	Description string `query:"description"`
	Year        int    `query:"year"`
	Page        int    `query:"page"`
	Limit       int    `query:"limit"`
}

type ReportUpdateRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	PeriodStart time.Time `json:"period_start" validate:"required"`
	PeriodEnd   time.Time `json:"period_end" validate:"required"`
	Year        int       `json:"year" validate:"required"`
	Quarter     *int      `json:"quarter"`
	Version     string    `json:"version" validate:"required"`
	ReportType  string    `json:"report_type" validate:"required"`
}

// response

type ReportResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
	Year        int       `json:"year"`
	Quarter     *int      `json:"quarter"`
	FileUrl     *string   `json:"fileurl"`
	Version     string    `json:"version"`
	Status      string    `json:"status"`
	UploadBy    string    `json:"upload_by"`
	ApprovedBy  *string   `json:"approved_by"`
	ReportType  string    `json:"report_type"`
}

type ReportPaginateResponse struct {
	Reports   []ReportResponse `json:"reports"`
	Page      int              `json:"page"`
	Limit     int              `json:"limit"`
	Total     int64            `json:"total"`
	TotalPage int64            `json:"total_page"`
}
