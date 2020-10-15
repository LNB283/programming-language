/////////////////exo2.go///////////////////////////
///objective(s):
///				- Create an empty map
///				- Print out the result
///				- Print out the type of the map
///				- Print out the lenght of the map
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	emptymap := map[string]int{}
	fmt.Printf("Empty Map:%v\n", emptymap)
	fmt.Printf("Type of empty map:%T\n", emptymap)
	fmt.Printf("Lenght:%d", len(emptymap))
}
