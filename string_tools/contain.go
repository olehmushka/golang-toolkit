package stringtools

import "strings"

func Contain(in, substr string) bool {
	return strings.Contains(strings.ToLower(in), strings.ToLower(substr))
}
