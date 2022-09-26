package pkg

type Error struct {
	StatusCode   int
	ErrorMessage string
	ErrorLog     string
}

func (e *Error) Error() string {
	return e.ErrorMessage
}

func NewError(statusCode int, message, log string) error {
	return &Error{
		StatusCode:   statusCode,
		ErrorMessage: message,
		ErrorLog:     log,
	}
}
