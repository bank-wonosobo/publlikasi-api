package request

type NewsCreateRequest struct {
	Title   string `json:"title" form:"title"`
	Slug    string `json:"slug" form:"slug"`
	Content string `json:"content" form:"content"`
	Author  string `json:"author" form:"author"`
	Image   string `json:"image" form:"image"`
}

type NewsGetQueryParams struct {
	Title   string `query:"title"`
	Content string `query:"content"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
}

type NewsUpdateRequest struct {
	Title   string  `json:"title" form:"title"`
	Slug    string  `json:"slug" form:"slug"`
	Content string  `json:"content" form:"content"`
	Author  string  `json:"author" form:"author"`
	Image   *string `json:"image" form:"image"`
}
