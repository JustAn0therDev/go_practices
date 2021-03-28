package calc

import (
	"errors"
	"log"
)

type Calculation struct {
	Operation string
	First int
	Second int
	Result int
}

func (c *Calculation) Calculate(numberOne int, numberTwo int, op *string) {
	var err error

	switch *op {
		case "sum":
			c.Result, err = sum(numberOne, numberTwo)
		case "sub":
			c.Result, err = sub(numberOne, numberTwo)
		case "mult":
			c.Result, err = mult(numberOne, numberTwo)
		case "div":
			c.Result, err = div(numberOne, numberTwo)
		default:
			err = errors.New("unknown operation")
	}

	if err != nil {
		log.Fatalln(err)
	}
}