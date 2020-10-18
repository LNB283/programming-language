/////////////////exo2.go///////////////////////////
///objective(s):
///				- Conver information from JSON format
///				- Create a structure, type person
///					-- This strucutre is the skeleton to format the information to json style
///					-- Use this site to create easily the skeleton of the structure person : https://mholt.github.io/json-to-go/
///				- Create a variable rawformat and assign the json raw format (use the result of the exercise1 for the json raw format)
///				- Create a variable ,type slice of bytes and you pass the first variable as parameter
///				- Create a loop to browse the content of the rawformat and print out the result
//////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"Age"`
	Height    int    `json:"Height"`
	Weight    int    `json:"Weight"`
	Hometown  string `json:"Hometown"`
	Country   string `json:"Country"`
	Retired   bool   `json:"Retired"`
}

func main() {

	rawformat := `[{"Firstname":"Rui","Lastname":"Hachimura","Age":22,"Height":203,"Weight":104,"Hometown":"Toyama","Country":"Japan","Retired":false},{"Firstname":"Michael","Lastname":"Jordan","Age":57,"Height":198,"Weight":98,"Hometown":"Brooklyn","Country":"USA","Retired":true},{"Firstname":"James","Lastname":"LeBron","Age":35,"Height":203,"Weight":113,"Hometown":"Akron","Country":"USA","Retired":false}]`
	bytevar := []byte(rawformat)

	fmt.Printf("%T\n", rawformat)
	fmt.Println(rawformat)
	fmt.Printf("%T\n", bytevar)
	var people []person

	//perform the convertion
	//Don't forget to manage potential errors
	errormgmt := json.Unmarshal(bytevar, &people)
	if errormgmt != nil {
		fmt.Printf("Error:%v", errormgmt)
	}
	fmt.Println("\nall of the data", people)
	for indice, value := range people {
		fmt.Printf("Person ID #%d\n", indice)
		fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", value.Firstname, value.Lastname, value.Age, value.Height, value.Weight, value.Hometown, value.Country, value.Retired)

	}
}
