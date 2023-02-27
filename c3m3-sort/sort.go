package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// sample input
//  8 923 93 7 634 32 83 37 9 8 42 12 83 63 0 83 12 7474

const goroutinesCount = 4 // number of goroutines

func sortNumbers(wg *sync.WaitGroup, numbers []int) {
	fmt.Println("sorting slice: ", numbers)
	sort.Ints(numbers)
	wg.Done()
}

func main() {
	fmt.Print("Input numbers split by space > ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	splitted := strings.Split(strings.TrimSpace(input), " ")
	numbers := make([]int, 0, len(splitted))
	for _, str := range splitted {
		num, _ := strconv.Atoi(str)
		numbers = append(numbers, num)
	}

	sliceSize := int(math.Max(float64(len(splitted)/goroutinesCount), 1))

	var waitGroup sync.WaitGroup

	for i := 0; i < goroutinesCount; i++ {
		from := int(math.Min(float64(i*sliceSize), float64(len(numbers))))
		to := int(math.Min(float64((i+1)*sliceSize), float64(len(numbers))))

		waitGroup.Add(1)
		go sortNumbers(&waitGroup, numbers[from:to])
	}

	waitGroup.Wait()
	sort.Ints(numbers)
	fmt.Println(numbers)
}
