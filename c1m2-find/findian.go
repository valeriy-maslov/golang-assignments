package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Provide input:")
	var input string
	fmt.Scan(&input)

	// Preparing string for analyze, trim spaces and make lower case for case-insensitive analyze
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	// Searching patterns...
	if strings.IndexRune(input, 'i') == 0 &&
		strings.ContainsRune(input, 'a') &&
		strings.LastIndex(input, "n") == len(input)-1 {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not found!\n")
	}
}
