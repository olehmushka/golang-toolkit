package numbertools

import (
	"sort"
	"strconv"
)

func ParseFloat64(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func ParseFloat32(str string) (float32, error) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return float32(f), nil
}

func Average(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}

	var total float64
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

func Median(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}

	sort.Float64s(x)

	half := len(x) / 2
	if isOdd(x) {
		return x[half]
	}

	return (x[half-1] + x[half]) / 2
}

func isOdd(x []float64) bool {
	return len(x)%2 != 0
}
