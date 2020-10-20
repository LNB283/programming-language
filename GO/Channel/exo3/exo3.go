/////////////////exo3.go///////////////////////////
///objective(s):
///				- Create a function
///					-- Param <- chan int
///					-- Print out all values on the channel
///				- Create a second fucntion
///					-- no param , return <- chan int
///					-- This function populates the channel by using un loop
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
