package main

import "fmt"

func main() {
	permutation([]int{1, 2, 3}, 0, 2)
}

func permutation(nums []int, left, right int) {
	if left == right {
		fmt.Println(nums)
	} else {
		for i := left; i <= right; i++ {
			// fmt.Println("i:", i, "nums:", nums, "left:", left, "right:", right)
			swap(nums, left, i)
			permutation(nums, left+1, right)
			swap(nums, left, i)
		}
	}
}

func swap(nums []int, i, j int) []int {
	nums[i], nums[j] = nums[j], nums[i]
	return nums
}
