package main

import "reflect"

// Use the sliding window method to find all anagrams in a string
func findAnagrams(full string, anagram string) []int {
	var (
		result []int
		sample = []rune(full)
		gram   = make(map[rune]int)
		mirror = make(map[rune]int)
	)

	for _, character := range anagram {
		gram[character]++
	}

	windowStart := 0
	windowLength := len(anagram)
	anagramStart := 0
	reset := true

	for index := 0; index < len(sample); index++ {
		if windowStart+windowLength == index {
			windowStart++
			index = windowStart
			mirror = make(map[rune]int)
			reset = true
		}

		character := sample[index]

		if _, ok := gram[character]; ok {
			if reset {
				anagramStart = index
			}
			mirror[character]++
			reset = false
		} else {
			reset = true
			windowStart = index
			if len(mirror) > 0 {
				mirror = make(map[rune]int)
			}
			continue
		}

		if reflect.DeepEqual(gram, mirror) {
			result = append(result, anagramStart)
		}
	}

	return result
}

/*
* Alternative solution
 */
func findAnagramsAlt(full string, anagram string) []int {
	var (
		result []int
		sample = []rune(full)
		gram   = make(map[rune]int)
		mirror = make(map[rune]int)
	)

	for _, character := range anagram {
		gram[character]++
		mirror[character]++
	}

	windowStart := 0
	windowLength := len(anagram)
	anagramStart := 0
	reset := true

	for index := 0; index < len(sample); index++ {
		if windowStart+windowLength == index {
			windowStart++
			index = windowStart
			mirror = copyMap(gram)
			reset = true
		}

		character := sample[index]

		if _, ok := gram[character]; ok {
			if reset {
				anagramStart = index
			}
			mirror[character]--
			if mirror[character] == 0 {
				delete(mirror, character)
			}
			reset = false
		} else {
			reset = true
			windowStart = index
			if len(mirror) > 0 {
				mirror = copyMap(gram)
			}
			continue
		}

		if len(mirror) == 0 {
			result = append(result, anagramStart)
		}
	}

	return result
}

func copyMap(input map[rune]int) map[rune]int {
	var result = make(map[rune]int, len(input))

	for key, value := range input {
		result[key] = value
	}

	return result
}
