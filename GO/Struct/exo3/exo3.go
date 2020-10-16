/////////////////exo3.go///////////////////////////
///objective(s):
///				- Example of anonymous strucuture
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	//Anonymous strucuture. We create a value person1
	person1 := struct {
		firstname string
		lastname  string
		age       int
		height    int
		weight    int
		hometown  string
		country   string
		retired   bool
	}{ //We assign information
		firstname: "Rui",
		lastname:  "Hachimura",
		age:       22,
		height:    203,
		weight:    104,
		hometown:  "Toyama",
		country:   "Japan",
		retired:   false,
	}

	//Print out the information for each value (person1 and person2)
	fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", person1.firstname, person1.lastname, person1.age, person1.height, person1.weight, person1.hometown, person1.country, person1.retired)

}
