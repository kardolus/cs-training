package main

import (
	"fmt"
	"math"
)

const (
	base  = 10
	prime = 113
)

// Great explanation of Rabin-Karp:
// https://www.cs.princeton.edu/courses/archive/spring14/cos226/lectures/53SubstringSearch.pdf

func hash(input string, length int) int { // Horner's method
	var hash int

	for index, char := range input {
		hash = (hash*base + int(char)) % prime
		if index == length-1 {
			break
		}
	}

	return hash
}

func findSubstring(text, pattern string) int {
	// preprocess
	patternHash := hash(pattern, len(pattern))
	textHash := hash(text, len(pattern))

	var runes = []rune(text)

	for i := len(pattern); i < len(text); i++ {
		if patternHash == textHash && pattern == text[i-len(pattern):i] {
			return i - len(pattern)
		}

		rollOff := int(runes[i-len(pattern)])
		rollOn := int(runes[i])
		correction := int(math.Pow10(len(pattern)-1)) % prime

		textHash = ((textHash+rollOff*(prime-correction))*base + rollOn) % prime
	}

	return -1
}

func main() {
	text := "this is the text to match against"
	pattern := "the"

	fmt.Println(findSubstring(text, pattern))
}
