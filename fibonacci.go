package main

import (
  "fmt"
)

func main() {  
  fmt.Println("Fibonacci(1)", FibonacciSlice(1))
  fmt.Println("Fibonacci(2)", FibonacciSlice(2))
  fmt.Println("Fibonacci(3)", FibonacciSlice(3))
  fmt.Println("Fibonacci(4)", FibonacciSlice(4))
  fmt.Println("Fibonacci(4)", FibonacciSlice(4))
  fmt.Println("Fibonacci(4)", FibonacciSlice(10))
  
  fmt.Println("Fibonacci(4)", FibonacciInPlace(10))
  
  fmt.Println("Fibonacci(4)", FibonacciRecursive(10))
}

func FibonacciSlice(index int) int { // O(N)
  if index <= 1 {
    return index
  }
  
  var sequence = []int{0, 1}
  
  for i := 2; i <= index; i++ {
    sequence = append(sequence, sequence[i - 1] + sequence[i - 2])
  }
  
  return sequence[index]
}

func FibonacciInPlace(index int) int { // O(N)
  if index <= 1 {
    return index
  }
  
  var ( 
    f1 = 0
    f2 = 1
  )
  
  for i := 2; i < index; i++ {
    tmp := f1
    f1 = f2
    f2 = f2 + tmp
  }
  
  return f1 + f2
}

func FibonacciRecursive(index int) int { // O(2^N) [branches^depth]
  if index <= 1 {
    return index
  }
  
  return FibonacciRecursive(index - 1) + FibonacciRecursive(index - 2)
}
