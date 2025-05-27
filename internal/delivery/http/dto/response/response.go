package response

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateReponseError(message string) Response[string] {
	return Response[string]{
		Message: message,
		Data:    "",
	}
}

func CreateReponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Message: "success",
		Data:    data,
	}
}

func CreatePaginateResponse[T any](data T, page int, limit int, total int64, totalPage int64) PaginateResponse[T] {
	return PaginateResponse[T]{
		Message:   "success",
		Data:      data,
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPage,
	}
}

type PaginateResponse[T any] struct {
	Message   string `json:"message"`
	Data      T      `json:"data"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Total     int64  `json:"total"`
	TotalPage int64  `json:"total_page"`
}
