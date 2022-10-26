package randomtools

import (
	"crypto/rand"
	"math/big"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

func RandIntInRange(min int, max int) (int, error) {
	if min >= max {
		return 0, wrapped_error.NewInternalServerError(nil, "[RandIntInRange] min >= max")
	}

	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(int64(min) + n.Int64()), nil
}

func RandInt(max int) (int, error) {
	return RandIntInRange(0, max)
}
