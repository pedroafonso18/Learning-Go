package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)
	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v", result, remainder)
	}
}

func printMe(printvalue string) {
	fmt.Println(printvalue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, nil
}
