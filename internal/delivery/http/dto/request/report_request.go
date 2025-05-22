package request

import "time"

type ReportCreateRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	PeriodStart time.Time `json:"period_start" validate:"required"`
	PeriodEnd   time.Time `json:"period_end" validate:"required"`
	Year        int       `json:"year" validate:"required"`
	Quarter     *int      `json:"quarter"`
	Version     string    `json:"version" validate:"required"`
	ReportType  string    `json:"report_type" validate:"required"`
}
