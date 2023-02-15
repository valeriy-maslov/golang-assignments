package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := make(sort.IntSlice, 0)
	for {
		var userInput string
		fmt.Println("Enter next number or X to exit:")
		fmt.Scan(&userInput)
		userInput = strings.ToLower(strings.TrimSpace(userInput))

		if strings.Compare("x", userInput) == 0 {
			break
		}

		number, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Wrong input, should be a number or X to exit")
		}

		numbers = append(numbers, number)
		sort.Sort(numbers)
		fmt.Println(numbers)
	}
}
