package maptools

import (
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

func PickOneByProb[T string](values map[T]float64) (T, error) {
	var zero T
	if len(values) == 0 {
		return zero, wrapped_error.NewInternalServerError(nil, "[GetRandomFromSeveral] Values count is zero")
	}

	var (
		valuesWithZeroCount int
		valuesWith1Count    int
	)
	for _, prob := range values {
		if prob == 0 {
			valuesWithZeroCount++
		}
		if prob == 1 {
			valuesWith1Count++
		}
	}
	if valuesWithZeroCount == len(values) {
		return zero, wrapped_error.NewInternalServerError(nil, "[GetRandomFromSeveral] All values are zero")
	}
	if valuesWith1Count > 1 {
		values = PrepareMapToPickRandomValue(values)
	}

	preparedValues := cloneMap(values)
	for value, prob := range preparedValues {
		preparedValues[value] = randomTools.PrepareProbability(prob - 0.01)
	}

	tempValues := make(map[T]float64)
	for {
		var iterValues map[T]float64
		if len(tempValues) == 0 {
			iterValues = cloneMap(preparedValues)
		} else {
			iterValues = cloneMap(tempValues)
		}
		tempValues = make(map[T]float64)

		for value, prob := range iterValues {
			if r, _ := randomTools.GetRandomBool(prob); r {
				tempValues[value] = prob
			}
		}

		if len(tempValues) == 1 {
			for value := range tempValues {
				return value, nil
			}
		}
	}
}

func PrepareMapToPickRandomValue[T string](values map[T]float64) map[T]float64 {
	var totalScore float64
	for _, score := range values {
		totalScore += score
	}

	out := make(map[T]float64, len(values))
	for key, value := range values {
		out[key] = value / totalScore
	}

	return out
}

func cloneMap[K string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64, V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](m map[K]V) map[K]V {
	clonedMap := make(map[K]V)
	for key, value := range m {
		clonedMap[key] = value
	}

	return clonedMap
}
