package calc

import (
	"errors"
)

func sub(numberOne int, numberTwo int) (int, error) {
	if numberOne == 0 || numberTwo == 0 {
		return 0, errors.New("both numbers to subtract must be provided")
	}

	return numberOne - numberTwo, nil
}