package wrapped_error

import (
	"errors"
	"testing"
)

func TestEquals(t *testing.T) {
	tCases := []struct {
		name        string
		left, right error
		expected    bool
	}{
		{
			name:     "should respond true for both nil operands",
			left:     nil,
			right:    nil,
			expected: true,
		},
		{
			name:     "should respond false for not nil left & nil right operands",
			left:     errors.New("abc"),
			right:    nil,
			expected: false,
		},
		{
			name:     "should respond false for not nil right & nil left operands",
			left:     nil,
			right:    errors.New("abc"),
			expected: false,
		},
		{
			name:     "should respond true for the same simple errors for both operands",
			left:     errors.New("abc"),
			right:    errors.New("abc"),
			expected: true,
		},
		{
			name:     "should respond false for different simple errors",
			left:     errors.New("abc"),
			right:    errors.New("abcd"),
			expected: false,
		},
		{
			name:     "should respond false for different wrapped errors",
			left:     NewInternalServerError(nil, "abc"),
			right:    NewInternalServerError(nil, "abcd"),
			expected: false,
		},
		{
			name:     "should respond true for the same wrapped errors for both operands",
			left:     NewInternalServerError(nil, "abc"),
			right:    NewInternalServerError(nil, "abc"),
			expected: true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if actual := Equals(tc.left, tc.right); actual != tc.expected {
				tt.Errorf("unexpectec result (expected=%t, actual=%t, left=%+v, right=%+v)", tc.expected, actual, tc.left, tc.right)
			}
		})
	}
}
