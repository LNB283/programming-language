/////////////////exo4.go///////////////////////////
///objective(s):
///				- Same s exo3.go but you use a select in the function to print out the content
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case value := <-c:
			fmt.Printf("%v\n", value)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}
