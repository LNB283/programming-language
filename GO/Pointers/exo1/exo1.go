/////////////////exo1.go///////////////////////////
///objective(s):
///				- Create a function :  switchvalues
///					--parameters : 2 pointers of integer
///				- Create 2 variables: pointera and pointerb and assign a value for each
///				- Print the initial informaiton (value and address)
///				- Call the function
///				- Print the updated information
//////////////////////////////////////////////////

package main

import (
	"fmt"
)

var pointera, pointerb int

func switchvalues(pointa *int, pointb *int) {

	temporarya := *pointa
	temporaryb := *pointb
	//switch the values
	*pointb = temporarya
	*pointa = temporaryb
	//fmt.Printf("Type of *pointa:%T\n", *pointa)

}

func main() {

	pointera := 23
	pointerb := 24

	fmt.Printf("Type of pointera:%T\tType of &pointera:%T\n", pointera, &pointera)
	fmt.Printf("Type of pointerb:%T\tType of &pointerb:%T\n", pointerb, &pointerb)
	fmt.Printf("Value of pointer a:%d\tAddress of pointer a:%v\n", pointera, &pointera)
	fmt.Printf("Value of pointer b:%d\tAddress of pointer b:%v\n", pointerb, &pointerb)

	switchvalues(&pointera, &pointerb)

	fmt.Printf("Update after switch functoin\n")
	//Output : values switch between pointer a and pointer b but used the same address
	fmt.Printf("Value of pointer a:%d\tAddress of pointer a:%v\n", pointera, &pointera)
	fmt.Printf("Value of pointer b:%d\tAddress of pointer b:%v\n", pointerb, &pointerb)
}
