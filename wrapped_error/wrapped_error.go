package wrapped_error

import "encoding/json"

type WrappedError struct {
	StatusCode  int
	Err         error
	Msg         string
	ExtendedErr *WrappedError
}

func New(statusCode int, err error, msg string) error {
	return &WrappedError{
		StatusCode:  statusCode,
		Msg:         msg,
		ExtendedErr: Cast(err),
	}
}

func (we *WrappedError) Error() string {
	if we == nil {
		return ""
	}

	var errMsg string
	if we.ExtendedErr != nil {
		errMsg = we.ExtendedErr.Error()
	}

	out := &outputMsg{
		Error:        errMsg,
		ErrorMessage: we.Msg,
	}

	errJSON, err := json.Marshal(out)
	if err != nil {
		return we.Msg
	}

	return string(errJSON)
}

func (we *WrappedError) ErrorMap() map[string]interface{} {
	var errMsg map[string]interface{}
	if we.ExtendedErr != nil {
		errMsg = we.ExtendedErr.ErrorMap()
	}

	return map[string]interface{}{
		"error":         errMsg,
		"error_message": we.Msg,
	}
}
