package main

import "fmt"

/*
	Example: decorating a function with some additional logging
*/

func doubleValue(value int) int {
	return value * 2
}

type doubleIt func(int) int

func decorator(input doubleIt) doubleIt {
	fmt.Println("Adding logging before")

	result := input

	fmt.Println("Adding logging after")

	return result
}

func main() {
	decoratedFN := decorator(doubleValue)
	fmt.Println("result:", decoratedFN(7))
}
