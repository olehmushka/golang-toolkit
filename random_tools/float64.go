package randomtools

import (
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	expRand "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

func RandFloat64InRange(min, max float64) (float64, error) {
	if min >= max {
		return 0, wrapped_error.NewInternalServerError(nil, "[RandFloat64InRange] min >= max")
	}

	rInt, err := RandIntInRange(0, 100)
	if err != nil {
		return 0, err
	}
	s := expRand.NewSource(uint64(rInt))
	return min + expRand.New(s).Float64()*(max-min), nil
}

func RandFloat64(max float64) (float64, error) {
	return RandFloat64InRange(0, max)
}

// RandFloat64Norm
// stdDev -standart deviation - σ^2; default = 1
// mean - μ (In probability theory, the expected value is a generalization of the weighted average.); default = 0
func RandFloat64Norm(stdDev, mean float64) (float64, error) {
	rInt, err := RandIntInRange(0, 100)
	if err != nil {
		return 0, err
	}
	s := expRand.NewSource(uint64(rInt))

	dist := distuv.Normal{
		Mu:    mean,   // Mean of the normal distribution
		Sigma: stdDev, // Standard deviation of the normal distribution
		Src:   s,
	}

	return dist.Rand(), nil
}

func RandFloat64NormInRange(min, max, stdDev, mean float64) (float64, error) {
	if min >= max {
		return 0, wrapped_error.NewInternalServerError(nil, "[RandFloat64NormInRange] min >= max")
	}

	return randFloat64NormInRange(min, max, stdDev, mean, 10)
}

func randFloat64NormInRange(min, max, stdDev, mean float64, count int) (float64, error) {
	var multiplier float64 = 10

	min *= multiplier
	max *= multiplier
	mean *= multiplier

	r, err := RandFloat64Norm(stdDev, mean)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not generate value by  norm distributed")
	}

	if r < min {
		if count == 0 {
			return min / multiplier, nil
		}
		return randFloat64NormInRange(min/multiplier, max/multiplier, stdDev, mean/multiplier, count-1)
	}
	if r > max {
		if count == 0 {
			return max / multiplier, nil
		}
		return randFloat64NormInRange(min/multiplier, max/multiplier, stdDev, mean/multiplier, count-1)
	}

	return r / multiplier, nil
}
