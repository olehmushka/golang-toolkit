package randomtools

import "github.com/olehmushka/golang-toolkit/wrapped_error"

const (
	asciiFirstLowcaseLetterCode = 97
	asciiLastLowcaseLetterCode  = 122
)

func RandString(len int) (string, error) {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		n, err := RandIntInRange(asciiFirstLowcaseLetterCode, asciiLastLowcaseLetterCode)
		if err != nil {
			return "", wrapped_error.NewInternalServerError(err, "can not random int value for string generation")
		}
		bytes[i] = byte(n)
	}

	return string(bytes), nil
}
