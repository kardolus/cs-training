package main

import (
  "fmt"
)

func main() {  
  var unsorted = []int{5, 4, 2, 8, 4, 7, 10, 1, 3}
  fmt.Println("unsorted:", unsorted)
  fmt.Println("sorted:", InsertionSort(unsorted))
}

// || = pivot
// 5,|| 4, 2, 8, 4, 7, 10, 1, 
// 4, 5,|| 2, 8, 4, 7, 10, 1, 
// 2, 4, 5,|| 8, 4, 7, 10, 1, 
// 2, 4, 5, 8,|| 4, 7, 10, 1, 

func InsertionSort(input []int) []int { // O(n^2)
  for pivot := 1; pivot < len(input); pivot++ {    
    for i := 0; i < pivot; i++ {
      if input[pivot] < input[i] {
        tmp := input[i]
        input[i] = input[pivot]
        input[pivot] = tmp
      } 
    }
  }
  
  return input
}
