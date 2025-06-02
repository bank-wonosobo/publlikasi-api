package dto

type AnnouncementGetQueryParams struct {
	Key            string `query:"key"`
	TargetAudience string `query:"target_audience"`
	IsActive       int    `query:"is_active"`
	Status         string `query:"status"`
	Page           int    `query:"page"`
	Limit          int    `query:"limit"`
}

type AnnouncementResponse struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	Content        string  `json:"content"`
	Author         string  `json:"author"`
	TargetAudience string  `json:"target_audience"`
	StartDate      string  `json:"start_date"`
	EndDate        string  `json:"end_date"`
	AttachmentUrl  *string `json:"attachment"`
	IsActive       bool    `json:"is_active"`
	Status         string  `json:"status"`
}
