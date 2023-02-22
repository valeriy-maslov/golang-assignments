package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (animal Cow) Eat() {
	fmt.Println("grass")
}

func (animal Cow) Move() {
	fmt.Println("walk")
}

func (animal Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (animal Bird) Eat() {
	fmt.Println("worms")
}

func (animal Bird) Move() {
	fmt.Println("fly")
}

func (animal Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (animal Snake) Eat() {
	fmt.Println("mice")
}

func (animal Snake) Move() {
	fmt.Println("slither")
}

func (animal Snake) Speak() {
	fmt.Println("hss")
}

func main() {
	animals := map[string]Animal{}

	for {
		var query, param1, param2 string
		fmt.Print("> ")
		fmt.Scanln(&query, &param1, &param2)

		if query == "newanimal" {
			switch param2 {
			case "cow":
				animals[param1] = Cow{}
			case "bird":
				animals[param1] = Bird{}
			case "snake":
				animals[param1] = Snake{}
			}
		} else if query == "query" {
			fmt.Print(param1, ": ")
			switch param2 {
			case "eat":
				animals[param1].Eat()
			case "move":
				animals[param1].Move()
			case "speak":
				animals[param1].Speak()
			}
		}
	}

}
