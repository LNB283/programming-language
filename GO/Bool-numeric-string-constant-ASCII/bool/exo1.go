/////////////////exo1.go///////////////////////////
///objective(s):
///				- create 4 variables
///					-- int1 is an integer
///					-- int2 is an integer
///					-- int3 is an integer with the same value as int1
///					-- string1 is a string
///					-- string2 is a string
///					-- string3 is a string with the same value as string1
///				- Perform multiple comparaison
///					-- use the list of operators: "!= / == / <= / >= / > / <"
///				- Print out the result
//////////////////////////////////////////////////
package main

import (
	"fmt"
)

func main() {

	int1 := 23
	int2 := 33
	int3 := 23
	string1 := "Hashimura"
	string2 := "Jordan"
	string3 := "Hashimura"

	fmt.Printf("Variable\t\tValue\t\tType\n")
	fmt.Printf("int1\t\t\t%d\t\t%T\nint2\t\t\t%d\t\t%T\nint3\t\t\t%d\t\t%T\nstring1\t\t\t%s\t%T\nstring2\t\t\t%s\t\t%T\nstring3\t\t\t%s\t%T\n\n", int1, int1, int2, int2, int3, int3, string1, string1, string2, string2, string3, string3)

	fmt.Println("int1 == int2 -->", int1 == int2)
	fmt.Println("int1 == int3 -->", int1 == int3)
	fmt.Println("string1 == string2 -->", string1 == string2) //Compare the 2 strings to determine if it is the same or not
	fmt.Println("string1 == string3 -->", string1 == string3)
	fmt.Printf("\n")
	fmt.Println("int1 != int2 -->", int1 != int2)
	fmt.Println("int1 != int3 -->", int1 != int3)
	fmt.Println("string1 != string2 -->", string1 != string2) //Compare the 2 strings to determine if they are different or not
	fmt.Println("string1 != string3 -->", string1 != string3)
	fmt.Printf("\n")
	fmt.Println("int1 >= int2 -->", int1 >= int2)
	fmt.Println("int1 >= int3 -->", int1 >= int3)
	fmt.Println("string1 >= string2 -->", string1 >= string2) // compare the alphabet range. H (position 9) J (position 12) so H is '<' than J
	fmt.Println("string1 >= string3 -->", string1 >= string3)
	fmt.Printf("\n")
	fmt.Println("int1 <= int2 -->", int1 <= int2)
	fmt.Println("int1 <= int3 -->", int1 <= int3)
	fmt.Println("string1 <= string2 -->", string1 <= string2)
	fmt.Println("string1 <= string3 -->", string1 <= string3)
	fmt.Printf("\n")
	fmt.Println("int1 < int2 -->", int1 < int2)
	fmt.Println("int1 < int3 -->", int1 < int3)
	fmt.Println("string1 < string2 -->", string1 < string2)
	fmt.Println("string1 < string3 -->", string1 < string3)
	fmt.Printf("\n")
	fmt.Println("int1 > int2 -->", int1 > int2)
	fmt.Println("int1 > int3 -->", int1 > int3)
	fmt.Println("string1 > string2 -->", string1 > string2)
	fmt.Println("string1 > string3 -->", string1 > string3)
	fmt.Printf("\n")
}
