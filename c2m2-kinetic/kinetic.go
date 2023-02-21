package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		// s = ½ a t2 + vot + so
		return 0.5*acc*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Println("Enter acceleration:")
	_, _ = fmt.Scan(&acceleration)
	fmt.Println("Enter initial velocity:")
	_, _ = fmt.Scan(&initialVelocity)
	fmt.Println("Enter initial displacement:")
	_, _ = fmt.Scan(&initialDisplacement)
	fmt.Println("Enter time:")
	_, _ = fmt.Scan(&time)

	displaceFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println(displaceFn(time))
}
