package exception

import "net/http"

type ExceptionErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ExceptionErr) Error() string {
	return r.Message
}

func newExecption(message, err string, code int, causes []Causes) *ExceptionErr {
	return &ExceptionErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func newBadRequestError(message string) *ExceptionErr {
	return &ExceptionErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadGateway,
	}
}

func newBadRequestValidationError(message string, causes []Causes) *ExceptionErr {
	return &ExceptionErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadGateway,
		Causes:  causes,
	}
}

func newInternalServerError(message string) *ExceptionErr {
	return &ExceptionErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func newNotFoundError(message string) *ExceptionErr {
	return &ExceptionErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func newForbiddenError(message string) *ExceptionErr {
	return &ExceptionErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
