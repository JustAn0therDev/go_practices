package calc

import (
	"errors"
)

func Div(numberOne int, numberTwo int) (int, error) {
	if numberOne == 0 || numberTwo == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return numberOne / numberTwo, nil
}