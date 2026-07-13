package apperror

import "net/http"

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return e.Msg
}

func BusinessError(msg string) *Error {
	return &Error{Code: http.StatusInternalServerError, Msg: msg}
}
