package error

// not found error
type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func NewNotFoundError(s string) *NotFoundError {
	return &NotFoundError{s}
}

// bad request error
type BadRequestError struct {
	Message string
}

func (b *BadRequestError) Error() string {
	return b.Message
}

func NewBadRequestError(s string) error {
	return &BadRequestError{s}
}

// internal server error
type ServerError struct {
	Message string
}

func (s *ServerError) Error() string {
	return s.Message
}

func NewServerError(s string) *ServerError {
	return &ServerError{s}
}

var NotFound *NotFoundError
var BadRequest *BadRequestError
var InternalServer *ServerError
