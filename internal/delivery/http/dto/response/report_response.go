package response

import "time"

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
