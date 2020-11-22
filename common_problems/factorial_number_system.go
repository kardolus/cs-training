package main

import "fmt"

func main() {
	elements := generateInput()

	perms := permutations(elements)
	for _, perm := range perms {
		fmt.Println(perm)
	}
}

// Using the Factorial Number System
// See also: https://en.wikipedia.org/wiki/Factorial_number_system
func permutations(input []string) [][]string {
	var result [][]string

	factorials := generateFactorials(len(input))
	for i := 0; i < factorials[len(input)]; i++ {
		var permutation []string

		tmp := make([]string, len(input))
		copy(tmp, input)

		positionCode := i
		for position := len(input); position > 0; position-- {
			selected := positionCode / factorials[position-1]
			permutation = append(permutation, tmp[selected])
			positionCode = positionCode % factorials[position-1]
			tmp = append(tmp[:selected], tmp[selected+1:]...)
		}

		result = append(result, permutation)
	}

	return result
}

func generateFactorials(size int) []int {
	result := []int{1}

	for i := 1; i <= size; i++ {
		result = append(result, result[i-1]*i)
	}

	return result
}

func generateInput() []string {
	return []string{"A", "B", "C", "D"}
	// return []string{"1", "2", "3"}
}
