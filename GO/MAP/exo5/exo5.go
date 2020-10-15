/////////////////exo5.go///////////////////////////
///objective(s):
///				- Create a map by using make function
///					-- Add 3 key/value pairs
///				- Print out the map
///				- Print out the lenght of the original map
///				- Update the value associated to User1
///				- Print out the new map
///				- Delete the key/value pair User3/33
///				- Print out the new map
///				- Print out the lenght of the updated map
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
	fmt.Printf("Original Map \n\t%v\n", newmap)
	fmt.Printf("Original map lenght : %v\n", len(newmap))
	newmap[`User1`] = 45
	fmt.Printf("User1 value updated \n\t%v\n", newmap)
	delete(newmap, `User3`)
	fmt.Printf("User3 deleted \n\t%v\n", newmap)
	fmt.Printf("Updated map lenght : %v\n", len(newmap))

}
