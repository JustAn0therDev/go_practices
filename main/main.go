package main

import (
	"errors"
	"fmt"
	"log"

	"example.com/calc"
)

type Calculation struct {
	operation string
	first int
	second int
	result int
}

func init() {
	log.SetFlags(0)
}

func main() {
	// Although Golang has pointers, there is no need to free allocated memory
	// for a pointer
	var calculation Calculation = Calculation{}
	fmt.Print("please select an operation (sum, sub, mult or div): ")
	fmt.Scanln(&calculation.operation)

	fmt.Print("insert the first value: ")
	fmt.Scanln(&calculation.first)

	fmt.Print("insert the second value: ")
	fmt.Scanln(&calculation.second)

	calculation.Calculate(calculation.first, calculation.second, &calculation.operation)

	fmt.Println(calculation.result)
}

func (c *Calculation) Calculate(numberOne int, numberTwo int, op *string) {
	var err error

	switch *op {
		case "sum":
			c.result, err = calc.Sum(numberOne, numberTwo)
		case "sub":
			c.result, err = calc.Sub(numberOne, numberTwo)
		case "mult":
			c.result, err = calc.Mult(numberOne, numberTwo)
		case "div":
			c.result, err = calc.Div(numberOne, numberTwo)
		default:
			err = errors.New("unknown operation")
	}

	if err != nil {
		log.Fatalln(err)
	}
}