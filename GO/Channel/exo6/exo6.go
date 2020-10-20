package main

import "fmt"

func main() {
	channel := make(chan int)

	go loopchannel(channel)
	for value := range channel {
		fmt.Println(value)
	}

	fmt.Println("about to exit")
}
func loopchannel(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}
