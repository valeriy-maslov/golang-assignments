package main

import (
	"fmt"
)

/*
*
A race condition or race hazard is a condition of software where the system's substantive behavior
depends on the sequence or timing of other uncontrollable events.
In my example, there are two goroutines simulating race conditions by changing the same value.
In the main function we try to print the changed value, but there is no guarantee it will be changed
because goroutines try communicating with that concurrently.
*/
var someValue = 0

func main() {

	go IncrementSomeValue()
	go IncrementSomeValue()

	fmt.Println(someValue)
}

func IncrementSomeValue() {
	someValue++
}
