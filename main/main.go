package main

import (
	"fmt"
	"log"

	"example.com/calc"
)

func main() {
	sum, err := calc.Sum(1, 2)

	if err != nil {
		log.Fatalln(err)
	}

	sub, err := calc.Sub(1, 2)

	if err != nil {
		log.Fatalln(err)
	}

	mult, err := calc.Mult(1, 2)

	if err != nil {
		log.Fatalln(err)
	}

	div, err := calc.Div(1, 2)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mult)
	fmt.Println(div)
}