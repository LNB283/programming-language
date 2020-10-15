/////////////////exo3.go///////////////////////////
///objective(s):
///				- Create a map by using make function
///					-- Add 3 key/value pairs
///				- Print out the result
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
	fmt.Println(newmap)

}
