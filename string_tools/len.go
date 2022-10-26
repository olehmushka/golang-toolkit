package stringtools

func GetLen(in string, def int) int {
	if l := len(in); l > 0 {
		return l
	}

	return def
}
