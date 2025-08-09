package dto

type Response[T any] struct {
	Code   int      `json:"code"`	
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponseError(message string) Response [string]{
	return Response[string]{
		Code : 400,
		Message: message,	
		Data: "",
	}
}

func CreateResponseErrorData(message string, data map[string]string) Response [string]{
	return Response[string]{
		Code : 422,
		Message: message,	
		Data: data,
	}
}

func CreateResponseSuccess[T any](message string, data T) Response [T]{
	return Response[T]{
		Code : 200,
		Message: message,	
		Data: data,
	}
}