/////////////////exo5.go///////////////////////////
///objective(s):
///				- Create an anonymous function and she's self-executed
///				- Use the comma ok idiom (it's like a boolean)
///				- Print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 24
	}()

	v, ok := <-c
	fmt.Printf("First iteration\nvalue:%v\tbool:%v\n", v, ok)

	close(c)

}
