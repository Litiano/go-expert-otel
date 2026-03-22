package http

type HttpError struct {
	Message string
	Code    int
}

func NewHttpError(message string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Code:    code,
	}
}
