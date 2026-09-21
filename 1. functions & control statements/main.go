package main

import (
	"errors"
	"fmt"
)

func main() {
	var value string = "Hello world"
	printMe(value)

	var numerator int = 16
	var denominator int = 12
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)
	var newResult, newRemainder = intDivisionAndRemainder(numerator, denominator)
	fmt.Printf("The result is %v & the remainder is %v \n", newResult, newRemainder)
	var eResult, eRemainder, err = divisionWithError(numerator, 15)
	if err != nil {
		fmt.Println(err.Error())
	} else if eRemainder == 0 {
		fmt.Printf("The result is %v", eResult)
	} else {
		fmt.Printf("The result is %v & the remainder is %v with error %v \n", eResult, eRemainder, err)
	}
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case eRemainder == 0:
		fmt.Printf("The result is %v", eResult)
	default:
		fmt.Printf("The result is %v & the remainder is %v with error %v \n", eResult, eRemainder, err)
	}
	switch eRemainder {
	case 0:
		fmt.Printf("The remainder was 0")
	case 1, 2:
		fmt.Printf("The remainder was niceeee, %v", eRemainder)
	default:
		fmt.Printf("The result is %v & the remainder is %v with error %v \n", eResult, eRemainder, err)
	}
}

func printMe(value string) {
	fmt.Println("Hello world!" + value)
}

func intDivision(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func intDivisionAndRemainder(numerator int, denominator int) (int, int) {
	var remainder int = numerator % denominator
	var result int = numerator / denominator
	return result, remainder
}

func divisionWithError(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var remainder int = numerator % denominator
	var result int = numerator / denominator
	return result, remainder, err
}
