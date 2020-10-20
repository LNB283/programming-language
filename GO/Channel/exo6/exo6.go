/////////////////exo6.go///////////////////////////
///objective(s):
///				- Create a function loopchannel
///					-- Populate the channel with xx values (here just 20)
///				- Print out the result
//////////////////////////////////////////////////
package main

import "fmt"

func main() {
	channel := make(chan int)

	go loopchannel(channel)
	for value := range channel {
		fmt.Println(value)
	}
}
func loopchannel(channel chan int) {
	for i := 0; i < 20; i++ {
		channel <- i
	}
	close(channel)
}
