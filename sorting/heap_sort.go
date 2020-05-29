package main

import (
	"fmt"
	"math"
)

func printHeap(array []int) {
	// row size is 2^n
	row := float64(0)
	prevRow := float64(0)

	for index := range array {
		fmt.Printf("[%d] ", array[index])
		if float64(index) == math.Pow(2, row)+prevRow-1 {
			fmt.Println()
			prevRow += math.Pow(2, row)
			row += 1.0
		}
	}
}

func height(array []int) int {
	var rowSize float64
	size := len(array)
	if size == 0 {
		return size
	}

	for i := float64(0); true; i += 1.0 {
		rowSize += math.Pow(2.0, i)
		if rowSize >= float64(size) {
			return int(i) + 1
		}
	}

	return 0
}

func maxHeapify(array []int, index int) {
	largest := index

	if leftChild(index) < len(array) && array[leftChild(index)] > array[largest] {
		largest = leftChild(index)
	}
	if rightChild(index) < len(array) && array[rightChild(index)] > array[largest] {
		largest = rightChild(index)
	}

	if largest != index {
		array[largest], array[index] = array[index], array[largest]
		maxHeapify(array, largest)
	}
}

func buildMaxHeap(array []int) {
	middle := len(array) / 2

	for i := middle; i >= 0; i-- {
		maxHeapify(array, i)
	}
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	return int(math.Round(float64(index)/2)) - 1
}

func main() {
	fmt.Println("Parent 5:", parent(5))
	fmt.Println("Parent 6:", parent(6))

	var unsorted = []int{5, 4, 2, 8, 4, 7, 10, 1, 3, 42, 42, 19, -16, 1000, -4, 6, 17}
	fmt.Println("unsorted:", unsorted)
	fmt.Println("PrintHeap >>>")
	printHeap(unsorted)
	fmt.Println("\n<<< PrintHeap")
	fmt.Println("height unsorted:", height(unsorted))
	fmt.Println("---> MaxHeapify(2)")
	maxHeapify(unsorted, 2)
	fmt.Println("PrintHeap >>>")
	printHeap(unsorted)
	fmt.Println("\n<<< PrintHeap")

	fmt.Println("---> BuildMaxHeap!!")
	buildMaxHeap(unsorted)
	fmt.Println("PrintHeap >>>")
	printHeap(unsorted)
	fmt.Println("\n<<< PrintHeap")
}
