package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Function to generate all permutations of cases
func generatePermutations(s string, index int, current string, results *[]string) {
	if index == len(s) {
		*results = append(*results, current)
		return
	}

	// Get the current character as a rune
	char := rune(s[index])

	if unicode.IsLetter(char) {
		// Generate permutations with the current character as lowercase
		generatePermutations(s, index+1, current+strings.ToLower(string(char)), results)

		// Generate permutations with the current character as uppercase
		generatePermutations(s, index+1, current+strings.ToUpper(string(char)), results)
	} else {
		// If it's not a letter, just continue with the same character
		generatePermutations(s, index+1, current+string(char), results)
	}
}

func main() {
	s := "fB&1 "
	var results []string
	generatePermutations(s, 0, "", &results)

	for _, perm := range results {
		fmt.Println(perm)
	}
}
