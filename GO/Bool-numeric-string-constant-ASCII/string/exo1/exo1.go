/////////////////exo1.go//////////////////////////////////////////////
///objective(s):
///				- Create a variable sentence
///					-- Type : String
///				- Assign the value and respect the form :
///										I Love
///										Golden
///										Labrador
///										And many "other dogs"
///				- Print out using raw string literals
///
///////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var sentence string
	//Assign the value to the variable sentence. Use ``
	sentence = `
	I Love
	Golden
	Labrador
	And many "other dogs"`
	fmt.Printf("With `` : \n%s\n", sentence)
}
