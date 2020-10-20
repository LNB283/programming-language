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

	v, ok = <-c
	fmt.Printf("Second iteration\nvalue:%v\t\tbool:%v\n", v, ok)

}
