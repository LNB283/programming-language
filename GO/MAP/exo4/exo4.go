/////////////////exo4.go///////////////////////////
///objective(s):
///				- Create a map by using make function
///					-- Add 3 key/value pairs
///				- Print out the map
///				- Add a new value
///				- Print out the new map
///				- Print out only the value associated to the new key/value pair created
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	newmap := make(map[string]int)
	newmap[`User1`] = 12
	newmap[`User2`] = 23
	newmap[`User3`] = 33
	fmt.Printf("Original map:\n%v\n", newmap)
	newmap[`User4`] = 45
	fmt.Printf("New map:\n%v\n", newmap)
	fmt.Printf("Value associated to Key User4 : %v", newmap[`User4`])
}
