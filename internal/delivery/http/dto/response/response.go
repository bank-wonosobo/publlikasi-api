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
