package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Println("Enter a name:")
	fmt.Scan(&name)
	fmt.Println("Enter an address:")
	fmt.Scan(&address)

	person := make(map[string]string)
	person["name"] = name
	person["address"] = address
	personJson, _ := json.Marshal(person)
	fmt.Println(string(personJson))

}
