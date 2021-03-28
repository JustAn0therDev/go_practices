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
	error error
}

func main() {
	// Although Golang has pointers, there is no need to free allocated memory
	// for a pointer
	var calculation Calculation = Calculation{}
	fmt.Println("please select an operation (sum, sub, mult or div): ")
	fmt.Scanln(&calculation.operation)

	fmt.Println("insert the first value: ")
	fmt.Scanln(&calculation.first)

	fmt.Println("insert the second value: ")
	fmt.Scanln(&calculation.second)

	calculation.result, calculation.error = callCalculation(calculation.first, calculation.second, &calculation.operation)

	if calculation.error != nil {
		log.Fatalln(calculation.error)
	}

	fmt.Println(calculation.result)
}

func callCalculation(numberOne int, numberTwo int, op *string) (int, error) {
	result := 0
	var err error

	switch *op {
		case "sum":
			result, err = calc.Sum(numberOne, numberTwo)
		case "sub":
			result, err = calc.Sub(numberOne, numberTwo)
		case "mult":
			result, err = calc.Mult(numberOne, numberTwo)
		case "div":
			result, err = calc.Div(numberOne, numberTwo)
		default:
			return 0, errors.New("unknown operation")
	}

	if err != nil {
		return 0, err
	}

	return result, nil
}