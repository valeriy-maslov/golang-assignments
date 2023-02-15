package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func main() {
	var filename string

	fmt.Println("Enter file name:")
	fmt.Scan(&filename)

	pwd, _ := os.Getwd()
	personsFileBytes, err := os.ReadFile(pwd + "/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	persons := make([]Person, 0)
	nameLines := strings.Split(string(personsFileBytes), "\n")
	for _, fullName := range nameLines {
		nameSplit := strings.Split(fullName, " ")
		persons = append(persons, Person{nameSplit[0], nameSplit[1]})
	}

	for _, person := range persons {
		fmt.Println(person)
	}
}
