/////////////////exo1.go///////////////////////////
///objective(s):
///				- create 6 variables:
///							-- Last Name (string)
///							-- First Name(string)
///							-- Age (int)
///							-- Height (int)
///							-- Weight (int)
///							-- Home Town (string)
///							-- Country (string)
///				- Print the result:
///							-- with Println
///							-- with Printf
//////////////////////////////////////////////////
package main

//Package fmt implements formatted I/O//
import (
	"fmt"
)

func main() {

	//Variable declaration and values assignement
	lastname := "Hachimura"
	firstname := "Rui"
	age := 22
	height := 203
	weight := 104
	hometown := "Toyama"
	country := "Japan"

	// Print result with Println
	fmt.Println("Print result with Println")
	fmt.Println(" Last Name: ", lastname, "\n", "First Name: ", firstname, "\n", "Age: ", age, "\n", "Height: ", height, "\n", "Weight: ", weight, "\n", "HomeTown: ", hometown, "\n", "Country: ", country)

	// Print result with Printf
	fmt.Printf("\nPrint result with Printf\n")
	fmt.Printf("Last Name:%s\nFirst Name:%s\nAge:%d\nHeight:%d\nWeight:%d\nHomeTown:%s\nCountry:%s", lastname, firstname, age, height, weight, hometown, country)
}
