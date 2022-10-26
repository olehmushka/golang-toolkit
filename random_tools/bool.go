package randomtools

func GetRandomBool(trueProb float64) (bool, error) {
	n, err := RandFloat64(1)
	if err != nil {
		return false, err
	}
	return n < trueProb, nil
}
