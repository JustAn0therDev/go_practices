package main

import (
	"fmt"
	"log"

	"example.com/calc"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// Although Golang has pointers, there is no need to free allocated memory for a pointer
	// Shorthand variable declarations can only be used in a function scope (or a structure like if, else, for...)
	calculation := calc.Calculation{}
	fmt.Print("please select an operation (sum, sub, mult or div): ")
	fmt.Scanln(&calculation.Operation)

	fmt.Print("insert the first value: ")
	fmt.Scanln(&calculation.First)

	fmt.Print("insert the second value: ")
	fmt.Scanln(&calculation.Second)

	calculation.Calculate(calculation.First, calculation.Second, &calculation.Operation)

	fmt.Println(calculation.Result)
}