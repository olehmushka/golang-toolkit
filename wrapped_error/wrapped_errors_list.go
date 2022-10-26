package wrapped_error

import "net/http"

func NewBadRequestError(err error, msg string) error {
	return New(http.StatusBadRequest, err, msg)
}

func NewUnauthorizedError(err error, msg string) error {
	return New(http.StatusUnauthorized, err, msg)
}

func NewForbiddenError(err error, msg string) error {
	return New(http.StatusForbidden, err, msg)
}

func NewNotFoundError(err error, msg string) error {
	return New(http.StatusNotFound, err, msg)
}

func NewInternalServerError(err error, msg string) error {
	return New(http.StatusInternalServerError, err, msg)
}
