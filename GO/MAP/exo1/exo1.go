/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create a map
///					-- Key is a string
///					-- value is a int
///					-- Add 3 pairs key/value
///				- Print out the result
///				-Print out the type of the map
///				- Print out the lenght of the map
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	var newmap = map[string]int{
		`User1`: 12,
		`User2`: 23,
		`User3`: 33}

	fmt.Printf("Map content:%v\n", newmap)
	fmt.Printf("Type:%T\n", newmap)
	fmt.Printf("Map lenght:%d\n", len(newmap))

}
