package main

import "fmt"

func MaximumCrossingArray(array []int, start, middle, end int) (startIndex, endIndex, maximum int) {
	maxUInt := ^uint(0)         // inverse of 0
	maxInt := int(maxUInt >> 1) // one int to the right (because first bit is used for the sign)
	minInt := -maxInt - 1       // zero is part of the possitive numbers

	var leftMax = minInt
	var rightMax = minInt

	var leftTotal, rightTotal int

	for i := middle; i >= 0; i-- {
		leftTotal += array[i]
		if leftTotal > leftMax {
			leftMax = leftTotal
			startIndex = i
		}
	}

	for i := middle + 1; i <= end; i++ {
		rightTotal += array[i]
		if rightTotal > rightMax {
			rightMax = rightTotal
			endIndex = i
		}
	}

	maximum = leftMax + rightMax

	return startIndex, endIndex, maximum
}

func MaximumSubarray(input []int, start, end int) (startIndex, endIndex, maximum int) {
	if start == end {
		return start, end, input[startIndex]
	}

	middle := (start + end) / 2

	leftStart, leftEnd, leftMax := MaximumSubarray(input, start, middle)
	rightStart, rightEnd, rightMax := MaximumSubarray(input, middle+1, end)
	crossStart, crossEnd, crossMax := MaximumCrossingArray(input, start, middle, end)

	if leftMax > rightMax && leftMax > crossMax {
		startIndex = leftStart
		endIndex = leftEnd
		maximum = leftMax
	} else if rightMax > leftMax && rightMax > crossMax {
		startIndex = rightStart
		endIndex = rightEnd
		maximum = rightMax
	} else {
		startIndex = crossStart
		endIndex = crossEnd
		maximum = crossMax
	}

	return startIndex, endIndex, maximum
}

func main() {
	stockPrices := []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}
	differences := make([]int, len(stockPrices)-1)

	for i := 1; i < len(stockPrices); i++ {
		differences[i-1] = stockPrices[i] - stockPrices[i-1]
	}

	fmt.Println("Stock Prices:", stockPrices)
	fmt.Println("Differences:", differences)
	fmt.Println("MaxCrossingArray Test:", wrap(MaximumCrossingArray(differences, 0, len(differences)/2, len(differences)-1)))
	fmt.Println("MaxSubArray:", wrap(MaximumSubarray(differences, 0, len(differences)-1)))

	var leetCodeArray = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("MaxSubArray:", wrap(MaximumSubarray(leetCodeArray, 0, len(leetCodeArray)-1)))

	var leftMax = []int{100, 200, 5, -10, -20, -20, -15, -5, -12, -18}
	fmt.Println("MaxSubArray:", wrap(MaximumSubarray(leftMax, 0, len(leftMax)-1)))
}

func wrap(input ...interface{}) []interface{} {
	return input
}
