package calc

import (
	"errors"
)

type Calculation struct {
	Operation string
	First int
	Second int
	result int
}

func (c *Calculation) Calculate(numberOne int, numberTwo int, op *string) (int, error) {
	var err error

	switch *op {
		case "sum":
			c.result, err = sum(numberOne, numberTwo)
		case "sub":
			c.result, err = sub(numberOne, numberTwo)
		case "mult":
			c.result, err = mult(numberOne, numberTwo)
		case "div":
			c.result, err = div(numberOne, numberTwo)
		default:
			err = errors.New("unknown operation")
	}

	if err != nil {
		return 0, err
	}

	return c.result, nil
}