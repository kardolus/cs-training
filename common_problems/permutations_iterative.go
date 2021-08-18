package main

import "fmt"

/*
  Use permutations of n-1 elements to generate permutations of n elements.
  Insert element n on each location of the n-1 permutation
  
  Note that there are n! permutations 
  
  Example:
  input: "abc"
  iteration 1 result: "a"
  iteration 2 result: "ab", "ba"
  iteration 3 result: "abc", "acb", "cab", "bac", "bca", cba"
  
  The method below is a waste of memory but it does work. Look into 
  Johnson and Trotter algorithm for the optimal way of doing this. 
*/

func main() {
  var result = permutations("abcd")
  fmt.Println(result)
  fmt.Println(len(result))
}

func permutations(input string) []string {
  result := []string{string(input[0])}
  
  for i := 1; i < len(input); i++ {
    var tmp []string
    
    for j := 0; j < len(result); j++ {
      tmp = append(tmp, insert(result[j], string(input[i]))...)
    }

    result = make([]string, len(tmp))
    copy(result, tmp)
  }
  
  return result
}

func insert(input string, char string) []string {
  var result []string
  
  for i := len(input); i >= 0; i-- {    
    result = append(result, input[:i]+char+input[i:])
  }
  
  return result
}
