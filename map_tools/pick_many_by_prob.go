package maptools

import "github.com/olehmushka/golang-toolkit/wrapped_error"

func PickManyByProb[T string](values map[T]float64, l int) ([]T, error) {
	result := make([]T, 0, l)
	for {
		val, err := PickOneByProb(values)
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pick several random from input")
		}
		if includes(result, val) {
			continue
		}
		result = append(result, val)
		if len(result) == l {
			break
		}
	}

	return result, nil
}

func includes[T string](values []T, v T) bool {
	for _, val := range values {
		if val == v {
			return true
		}
	}

	return false
}
