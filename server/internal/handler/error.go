package handler

type HandlerError struct {
	code int
	msg  string
}

func (e HandlerError) Error() string {
	return e.msg
}

func NewError(code int, msg string) *HandlerError {
	return &HandlerError{
		code,
		msg,
	}
}

func (e HandlerError) Code() int {
	return e.code
}
