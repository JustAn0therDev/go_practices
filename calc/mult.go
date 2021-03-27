package calc

import (
	"errors"
)

func Mult(numberOne int, numberTwo int) (int, error) {
	if numberOne == 0 || numberTwo == 0 {
		return 0, errors.New("both numbers to multiply must be provided")
	}

	return numberOne * numberTwo, nil
}