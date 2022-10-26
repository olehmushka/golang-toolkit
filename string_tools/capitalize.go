package stringtools

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(in string) string {
	if in == "" {
		return ""
	}

	return cases.Title(language.English).String(in)
}
