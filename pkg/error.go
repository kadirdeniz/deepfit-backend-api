package pkg

type Error struct {
	StatusCode   int
	ErrorMessage string
	ErrorInfo    error
}

func (e *Error) Error() string {
	return e.ErrorMessage
}

func NewError(statusCode int, message string, info error) error {
	return &Error{
		StatusCode:   statusCode,
		ErrorMessage: message,
		ErrorInfo:    info,
	}
}
