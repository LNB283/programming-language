/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create a function :  sum
///					-- Parameters : veriadic param of int
///					-- Return : int
///				- create a second function on the same package (in a different file : main_test.go)
///				- This second function should be a table of test
///					-- Name : TestSum
///					-- Parameters : (t *testing.T)
///						---Create a structure test
///							: testdataslice is a slice of integer
///							: testdataresult is a integer
///						---Create a variable testdata and you assign a list of values
///					--create a loop to perform the test
//////////////////////////////////////////////////

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
	//return sum + 2
}
