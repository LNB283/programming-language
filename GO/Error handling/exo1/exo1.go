/////////////////exo1.go///////////////////////////
///objective(s):
///				- Reuse the exercise: JSON/exo1/exo1.go
///				- Update the error handling part by using "log.Println"
//////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
	Height    int
	Weight    int
	Hometown  string
	Country   string
	Retired   bool
}

func main() {

	//Create the value person1, Type: person
	// Be careful : For each variable, use Uppercase on the first letter!!!
	//If you don't do that, your slice doesn't contain the information
	person1 := person{
		Firstname: "Rui",
		Lastname:  "Hachimura",
		Age:       22,
		Height:    203,
		Weight:    104,
		Hometown:  "Toyama",
		Country:   "Japan",
		Retired:   false,
	}
	//Create the value person2, Type: person
	person2 := person{
		Firstname: "Michael",
		Lastname:  "Jordan",
		Age:       57,
		Height:    198,
		Weight:    98,
		Hometown:  "Brooklyn",
		Country:   "USA",
		Retired:   true,
	}

	person3 := person{
		Firstname: "James",
		Lastname:  "LeBron",
		Age:       35,
		Height:    203,
		Weight:    113,
		Hometown:  "Akron",
		Country:   "USA",
		Retired:   false,
	}

	nbaplayers := []person{person1, person2, person3}

	fmt.Println(nbaplayers)
	//Error handling
	jsonformat, err := json.Marshal(nbaplayers)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(jsonformat))
}
