package main

import (
	"fmt"
)

const (
	Smaller = iota - 1
	Equal
	Larger
)

func BinarySearch(sorted []int, value int) int {
	if len(sorted) == 0 {
		return -1
	}

	switch compareMiddle(sorted, value) {
	case Equal:
		return 0
	case Larger:
		return BinarySearch(sorted[len(sorted)/2+1:], value)
	case Smaller:
		return BinarySearch(sorted[0:len(sorted)/2], value)
	}

	return -1
}

func compareMiddle(array []int, value int) int {
	middle := len(array) / 2

	if value == array[middle] {
		return Equal
	} else if value > array[middle] {
		return Larger
	}

	return Smaller
}

func main() {
	var sorted = []int{-10, -8, -5, 0, 4, 16, 18, 23, 26, 30, 35, 41, 50, 55, 59, 67, 69}
	fmt.Println("compareMiddle(20)", compareMiddle(sorted, 20))
	fmt.Println("BinarySearch(41)", BinarySearch(sorted, 41))
	fmt.Println("BinarySearch(41)", BinarySearch(sorted, 3))
}
