package wrapped_error

import "errors"

func Equals(left, right error) bool {
	if left == nil || right == nil {
		return errors.Is(left, right)
	}

	return left.Error() == right.Error()
}
