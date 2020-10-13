/////////////////exo5.go///////////////////////////
///objective(s):
///				- Use the code from exo4.go
///				- create another variable (outside the main function --> Package level)
///							-- convertjerseynumber is an integer
///				- In the main function
///							-- convert jnumb (type jerseynumber) to integer
///							-- assign this value to convertjerseynumber
///				- Print the Type of jerseynumber
///				- Print the value of jnum
///				- Print the Type of convertjerseynumber
///				- Print the value of convertjerseynumber
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

type jerseynumber int

var convertjerseynumber int

func main() {

	//Variable declaration
	var lastname string
	var firstname string
	var age int
	var height int
	var weight int
	var hometown string
	var country string
	var righthand bool
	var lefthand bool
	// create jnumb variable with the new type jerseynumber
	var jnum jerseynumber
	//Variable declaration and values assignement
	lastname = "Hachimura"
	firstname = "Rui"
	age = 22
	height = 203
	weight = 104
	hometown = "Toyama"
	country = "Japan"
	righthand = true
	// Assign a value to jnum
	jnum = 8
	//Print the result
	fmt.Printf("Last Name:%s\nFirst Name:%s\nAge:%d\nHeight:%d\nWeight:%d\nHomeTown:%s\nCountry:%s\nRight Hand:%t\nLeft Hand:%t\nJersey Number:%d\n", lastname, firstname, age, height, weight, hometown, country, righthand, lefthand, jnum)
	fmt.Printf("\n")
	//Print out the value and type of the variable jnum
	fmt.Printf("Value of Jersey Number is:%d\n", jnum)
	fmt.Printf("Type of Jersey Number is:%T\n", jnum)
	//Conversion of jnum
	convertjerseynumber = int(jnum)
	//Print out the value and type of the variable convertjerseynumber
	fmt.Printf("Value of convertjerseynumber is:%d\n", convertjerseynumber)
	fmt.Printf("Type of convertjerseynumber is:%T\n", convertjerseynumber)
}
