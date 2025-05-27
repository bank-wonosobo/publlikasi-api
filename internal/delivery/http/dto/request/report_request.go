package request

import "time"

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
