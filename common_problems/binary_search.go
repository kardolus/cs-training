package main

import (
	"fmt"
)

const (
	Smaller = iota - 1
	Equal
	Larger
)

// Here's iterative. 
func search(nums []int, target int) int {
    start := 0
    end := len(nums)
    
    for {
        if start == end - 1 {
            if nums[start] == target {
                return 0
            } 
            return -1
        }
        
        middle := (end+start)/2
        if nums[middle] == target {
            return middle
        }
        if nums[middle] < target {
            start = middle
            continue
        }
        end = middle
    }
}

// Here's recursive
func BinarySearch(sorted []int, start, end, value int) int {
	if start == end-1 {
		if sorted[start] == value {
			return start
		}
		return -1
	}

	switch compareMiddle(sorted, start, end, value) {
	case Equal:
		return (start + end) / 2
	case Larger:
		return BinarySearch(sorted, (start+end)/2, end, value)
	case Smaller:
		return BinarySearch(sorted, start, (start+end)/2, value)
	}

	return -1
}

func compareMiddle(array []int, start, end, value int) int {
	middle := (start + end) / 2

	if value == array[middle] {
		return Equal
	} else if value > array[middle] {
		return Larger
	}

	return Smaller
}

func main() {
	var sorted = []int{-10, -8, -5, 0, 4, 16, 18, 23, 26, 30, 35, 41, 50, 55, 59, 67, 69}
	fmt.Println("compareMiddle(20)", compareMiddle(sorted, 0, len(sorted), 20))
	fmt.Println("BinarySearch(41)", BinarySearch(sorted, 0, len(sorted), 41))
	fmt.Println("BinarySearch(3)", BinarySearch(sorted, 0, len(sorted), 3))
	fmt.Println("BinarySearch(-10)", BinarySearch(sorted, 0, len(sorted), -10))
	fmt.Println("BinarySearch(69)", BinarySearch(sorted, 0, len(sorted), 69))
	fmt.Println("BinarySearch(67)", BinarySearch(sorted, 0, len(sorted), 67))
}
