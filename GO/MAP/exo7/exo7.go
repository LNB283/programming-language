////////////////exo7.go///////////////////////////
///objective(s):
///				- Create a map by using make function
///					-- Add 5 key/value pairs
///				- Print out the map with the format
///					-- key: Userx value: y
///				- Print out the lenght
///				- Clear the map
///				- Print out the map and the lenght
//////////////////////////////////////////////////
package main

import "fmt"

func main() {

	newmap := make(map[string]int)
	newmap[`User1`] = 12
	newmap[`User2`] = 23
	newmap[`User3`] = 33
	newmap[`User4`] = 24
	newmap[`User5`] = 45

	for indice, value := range newmap {
		fmt.Printf("Key:%v\tValue:%v\n", indice, value)
	}
	fmt.Printf("Map lenght: %d\n", len(newmap))
	newmap = make(map[string]int)
	fmt.Printf("Map after cleaning %v\n", newmap)
	fmt.Printf("Map lenght: %d\n", len(newmap))

}
