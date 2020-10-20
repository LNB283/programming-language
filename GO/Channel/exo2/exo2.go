/////////////////exo2.go///////////////////////////
///objective(s):
///				- Create bidirectional channel with buffer
///				- Print the type
///				- Print the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	// channel type chan int with buffer equal to 1 means you pass only 1 value
	channel := make(chan int, 1)
	// Put value on channel
	channel <- 24
	//Print out the value passed in the channel
	fmt.Println(<-channel)
	//Print the Type
	fmt.Printf("Type:%T", channel)

}
