package main

import (
	"fmt"
)

func main() {

	fmt.Println("1 + 1 =", sum(1, 1))
	fmt.Println("2 + 2 =", sum(2, 2))
	fmt.Println("3 + 3 =", sum(3, 3))

}

func sum(value ...int) int {

	sum := 0
	for _, val := range value {
		sum = sum + val
	}
	return sum
	//Test if our test function works good and raise the error flag
	//return sum +2
}
