package randomtools

import (
	goNanoID "github.com/matoous/go-nanoid"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

func NanoString(length int) (string, error) {
	id, err := goNanoID.ID(length)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate nano id")
	}

	return id, nil
}
