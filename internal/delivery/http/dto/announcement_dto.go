package dto

type AnnouncementGetQueryParams struct {
	Key            string `query:"key"`
	TargetAudience string `query:"target_audience"`
	IsActive       int    `query:"is_active"`
	Status         string `query:"status"`
	Page           int    `query:"page"`
	Limit          int    `query:"limit"`
}

type AnnouncementCreateReq struct {
	Title          string  `json:"title" form:"title" validate:"required"`
	Content        string  `json:"content" form:"content" validate:"required"`
	Author         string  `json:"author" form:"author" validate:"required"`
	TargetAudience string  `json:"target_audience" form:"target_audience" validate:"required"`
	StartDate      string  `json:"start_date" form:"start_date" validate:"required"`
	EndDate        string  `json:"end_date" form:"end_date" validate:"required"`
	Attachment     *string `json:"attachment" form:"attachment"`
}

type AnnouncementResponse struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	Content        string  `json:"content"`
	Author         string  `json:"author"`
	TargetAudience string  `json:"target_audience"`
	StartDate      string  `json:"start_date"`
	EndDate        string  `json:"end_date"`
	AttachmentUrl  *string `json:"attachment_url"`
	IsActive       bool    `json:"is_active"`
	Status         string  `json:"status"`
}
