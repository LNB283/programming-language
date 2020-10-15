/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create a structure person
///							-- Last Name (string)
///							-- First Name(string)
///							-- Age (int)
///							-- Height (int)
///							-- Weight (int)
///							-- Home Town (string)
///							-- Country (string)
///				- Create 2  values of type Person
///				- Assign a value for each value of type person
///				- Print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

//Create the structure person
type person struct {
	firstname string
	lastname  string
	age       int
	height    int
	weight    int
	hometown  string
	country   string
	retired   bool
}

func main() {

	//Create the value person1, Type: person
	person1 := person{
		firstname: "Rui",
		lastname:  "Hachimura",
		age:       22,
		height:    203,
		weight:    104,
		hometown:  "Toyama",
		country:   "Japan",
		retired:   false,
	}
	//Create the value person2, Type: person
	person2 := person{
		firstname: "Michael",
		lastname:  "Jordan",
		age:       57,
		height:    198,
		weight:    98,
		hometown:  "Brooklyn",
		country:   "USA",
		retired:   true,
	}
	//Print out the information for each value (person1 and person2)
	fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", person1.firstname, person1.lastname, person1.age, person1.height, person1.weight, person1.hometown, person1.country, person1.retired)
	fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", person2.firstname, person2.lastname, person2.age, person2.height, person2.weight, person2.hometown, person2.country, person2.retired)

}
