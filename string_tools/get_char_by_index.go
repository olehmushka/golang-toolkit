package stringtools

func GetCharByIndex(in string, index int, def string) string {
	if len(in) <= index {
		return def
	}

	for i, char := range in {
		if i == index {
			return string(char)
		}
	}

	return def
}
