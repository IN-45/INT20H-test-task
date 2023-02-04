package customerrors

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *AlreadyExistsError {
	return &AlreadyExistsError{
		message: message,
	}
}

func (e *NotFoundError) Error() string {
	return e.message
}
