package main

import (
	"fmt"
)

// Partition uses the last element as the initial pivot and returns the final position
// of the pivot
func Partition(input []int) int {
	i := 0
	for j := 0; j < len(input); j++ {
		if input[j] < input[len(input)-1] {
			input[i], input[j] = input[j], input[i]
			i++
		}
	}
	input[i], input[len(input)-1] = input[len(input)-1], input[i]

	return i
}

func QuickSort(input []int) {
	if len(input) <= 1 {
		return
	}

	pivot := Partition(input)
	QuickSort(input[:pivot])
	QuickSort(input[pivot+1:])
}

func main() {
	var unsorted = []int{5, 4, 2, 101, 18, 3, 700, -15, 8, 4, 7, 10, 1, 3}
	fmt.Println("unsorted:", unsorted)
	fmt.Println("Partition(unsorted, 0, len(unsorted) - 1):", Partition(unsorted))
	fmt.Println("array after 1 partition:", unsorted)
	QuickSort(unsorted)
	fmt.Println("sorted:", unsorted)
}
