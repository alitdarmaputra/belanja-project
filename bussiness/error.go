package bussiness

type UnauthorizedError struct {
	message string
}

func (error UnauthorizedError) Error() string {
	return error.message
}

func NewUnauthorizedError(error string) error {
	return &UnauthorizedError{
		message: error,
	}
}

type NotFoundError struct {
	message string
}

func (error NotFoundError) Error() string {
	return error.message
}

func NewNotFoundError(error string) error {
	return &NotFoundError{
		message: error,
	}
}

type DuplicateEntryError struct {
	message string
}

func (error *DuplicateEntryError) Error() string {
	return error.message
}

func NewDuplicateEntryError(error string) error {
	return &DuplicateEntryError{
		message: error,
	}
}
