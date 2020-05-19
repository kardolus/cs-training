package main

import (
	"fmt"
)

func main() {
	var unsorted = []int{5, 4, 2, 8, 4, 7, 10, 1, 3}
	fmt.Println("unsorted:", unsorted)
	fmt.Println("sorted:", BubbleSort(unsorted))
}

func BubbleSort(input []int) []int { // O(n^2)
	for i := 0; i < len(input)-1; i++ {
		bubble(input, len(input)-1-i)
	}

	return input
}

func bubble(input []int, index int) { // O(n)
	for i := 0; i < index; i++ {
		if input[i] > input[i+1] && i < len(input)-1 {
			input[i], input[i+1] = input[i+1], input[i]
		}
	}
}
