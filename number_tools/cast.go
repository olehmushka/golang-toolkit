package numbertools

import (
	"fmt"
	"strconv"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

func Float64ToBytes(in float64) []byte {
	return []byte(Float64ToString(in))
}

func BytesToFloat64(in []byte) (float64, error) {
	return StringToFloat64(string(in))
}

func Float64ToString(in float64) string {
	return fmt.Sprint(in)
}

func StringToFloat64(in string) (float64, error) {
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not convert float64 to string")
	}

	return out, nil
}

func StringToInt(in string) (int, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not convert string to int")
	}

	return out, nil
}
