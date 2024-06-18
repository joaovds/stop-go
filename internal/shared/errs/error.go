package errs

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewError(message string) *Error {
	defaultStatus := 500

	return &Error{
		Message: message,
		Status:  defaultStatus,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) SetStatus(status int) *Error {
	e.Status = status
	return e
}

func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

func (e *Error) IsError() bool {
	return e != nil
}

func (e *Error) IsNil() bool {
	return e == nil
}
