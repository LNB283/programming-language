////////////////exo6.go///////////////////////////
///objective(s):
///				- Create a map by using make function
///					-- Add 5 key/value pairs
///				- Print out the map with the format
///					-- key: Userx value: y
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

}
