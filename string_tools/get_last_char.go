package stringtools

func GetLastStrChar(in string) string {
	if len(in) == 0 {
		return ""
	}
	if len(in) == 1 {
		return in
	}

	return string(in[len(in)-1])
}
