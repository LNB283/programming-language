/////////////////exo2.go///////////////////////////
///objective(s):
///				- Create a structure person
///							-- Last Name (string)
///							-- First Name(string)
///							-- Age (int)
///							-- Height (int)
///							-- Weight (int)
///							-- Home Town (string)
///							-- Country (string)
///				- Create a seecond structure : nbaactivity (embded structure)
///							-- person
///							-- Retired (booolean)
///				- Create 2  values of type Person
///				- Assign a value for each value of type person
///				- Print out the result
//////////////////////////////////////////////////

package main

import (
	"fmt"
)

//Create the structure person
//Personn is defined by generic information and nbaactivity
type person struct {
	firstname string
	lastname  string
	age       int
	height    int
	weight    int
	hometown  string
	country   string
	nbaactivity
}

//Create nbaacitivity to determined if a person is retired or not in this business
type nbaactivity struct {
	retired bool
}

func main() {

	//Create the value person1, Type: person
	person1 := person{
		nbaactivity: nbaactivity{
			retired: false,
		},
		firstname: "Rui",
		lastname:  "Hachimura",
		age:       22,
		height:    203,
		weight:    104,
		hometown:  "Toyama",
		country:   "Japan",
	}
	//Create the value person2, Type: person

	person2 := person{
		nbaactivity: nbaactivity{
			retired: true,
		},
		firstname: "Michael",
		lastname:  "Jordan",
		age:       57,
		height:    198,
		weight:    98,
		hometown:  "Brooklyn",
		country:   "USA",
	}

	person3 := person{
		nbaactivity: nbaactivity{
			retired: false,
		},
		firstname: "James",
		lastname:  "LeBron",
		age:       35,
		height:    203,
		weight:    113,
		hometown:  "Akron",
		country:   "USA",
	}

	//Print out the information for each value
	fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", person1.firstname, person1.lastname, person1.age, person1.height, person1.weight, person1.hometown, person1.country, person1.retired)
	fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", person2.firstname, person2.lastname, person2.age, person2.height, person2.weight, person2.hometown, person2.country, person2.retired)
	fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", person3.firstname, person3.lastname, person3.age, person3.height, person3.weight, person3.hometown, person3.country, person3.retired)

}
