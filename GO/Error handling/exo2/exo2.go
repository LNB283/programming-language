/////////////////exo2.go///////////////////////////
///objective(s):
///				- Reuse the exercise: exo1.go
///				- Create a function ConverttoJSON
///					-- Parameters: interface{}
///					-- Return slice of byte and error
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

//Convert ot JSON format
func converToJSON(varinterface interface{}) ([]byte, error) {
	vartemp, err := json.Marshal(varinterface)
	if err != nil {
		//Use the function Errorf to print out the error
		return []byte{}, fmt.Errorf("Try to convert to Matshal form the information: %v\nError in the function converToJSO :%v", varinterface, err)
	}
	return vartemp, nil
}

func main() {

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
	jsonformat, err := converToJSON(nbaplayers)
	if err != nil {
		//Use Fatalln to print out the information concerning a Fatal error (means error that can block the progress of the code)
		log.Fatalln(err)
	}

	fmt.Printf("%v\n", string(jsonformat))
}
