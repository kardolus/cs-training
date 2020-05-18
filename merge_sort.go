package main

import (
	"fmt"
)

func main() {
	fmt.Println("merged:", merge([]int{1, 2, 4, 8}, []int{2, 5, 10, 89}))
	fmt.Println("merged:", merge([]int{1, 2, 4}, []int{2, 5, 10, 89}))

	var unsorted = []int{14, 5, 4, 1, 16, 18, 20, 2, 4, 6, 7}

	fmt.Println("--- MergeSort ---")
	fmt.Println("unsorted:", unsorted)
	fmt.Println("sorted:", MergeSort(unsorted))
}

func MergeSort(input []int) []int {
	fmt.Printf("MergeSort(%v)\n", input)
	if len(input) == 1 {
		return input
	}

	middle := len(input) / 2

	var (
		left  = make([]int, middle)
		right = make([]int, len(input)-middle)
	)

	for i := range input {
		if i < middle {
			left[i] = input[i]
			continue
		}
		right[i-middle] = input[i]
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	fmt.Printf("merge(%v, %v)\n", left, right)
	var (
		result     = make([]int, len(left)+len(right))
		leftIndex  = 0
		rightIndex = 0
	)

	for i := 0; i < len(result); i++ {
		if len(left) == leftIndex {
			result[i] = right[rightIndex]
			rightIndex++
			continue
		}
		if len(right) == rightIndex {
			result[i] = left[leftIndex]
			leftIndex++
			continue
		}
		if left[leftIndex] <= right[rightIndex] {
			result[i] = left[leftIndex]
			leftIndex++
			continue
		}
		result[i] = right[rightIndex]
		rightIndex++
	}

	return result
}
