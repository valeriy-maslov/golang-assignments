package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		var animal, action string
		fmt.Println("Choose the animal (cow, bird or snake) or Ctrl+C to exit:")
		fmt.Print("> ")
		fmt.Scan(&animal)
		fmt.Println("Choose the action (eat, move or speak) or Ctrl+C to exit:")
		fmt.Print("> ")
		fmt.Scan(&action)

		animal = strings.TrimSpace(animal)
		action = strings.TrimSpace(action)

		chosenAnimal := animals[animal]
		switch action {
		case "eat":
			chosenAnimal.Eat()
		case "move":
			chosenAnimal.Move()
		case "speak":
			chosenAnimal.Speak()
		}
	}

}
