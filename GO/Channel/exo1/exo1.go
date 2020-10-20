/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create a function
///						-- with parameter chan int
///						-- assign a value on the channel
///				- Create a variable channel by using make
///				- Create a routine
///				- Print the type
///				- Print the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	channel := make(chan int)
	go testchannel(channel)
	fmt.Printf("Value on the channel:%v\n", <-channel)
	fmt.Printf("Type:%T\n", channel)

}

func testchannel(channel chan int) {
	channel <- 23
}
