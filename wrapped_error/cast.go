package wrapped_error

import (
	"errors"
	"net/http"
)

func Cast(err error) *WrappedError {
	if err == nil {
		return nil
	}

	var wErr *WrappedError
	if errors.As(err, &wErr) {
		return wErr
	}

	return &WrappedError{
		StatusCode: http.StatusInternalServerError,
		Err:        err,
		Msg:        "",
	}
}
