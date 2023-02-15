package main

import "fmt"

func main() {
	numbers := make([]int, 10)

	for idx := 0; idx < 10; idx++ {
		var nextNumber int
		fmt.Println("Enter next number:")
		_, err := fmt.Scan(&nextNumber)
		if err != nil {
			fmt.Println(err)
		} else {
			numbers[idx] = nextNumber
		}
	}

	BubbleSort(numbers)
	fmt.Println(numbers)
}

func Swap(numbers []int, idx int) {
	temp := numbers[idx]
	numbers[idx] = numbers[idx+1]
	numbers[idx+1] = temp
}

func BubbleSort(numbers []int) {
	var i, j int
	for i = 0; i < len(numbers)-1; i++ {
		for j = 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}
