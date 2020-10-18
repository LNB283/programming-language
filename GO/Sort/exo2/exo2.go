/////////////////exo2.go///////////////////////////
///objective(s):
///				- Sort person based on their Weight
///				- Create sortbyheight, type []person
///				- create 3 functions
///					-- first func: return the lenght of the slice
///					-- second func : switch the position based on the height value
///					-- third funciton : perform the comparaison between the Height
///				- Print out the result
///			Bonus:
///				- Use the same pattern to sort people by their Age
//////////////////////////////////////////////////
package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstname string
	lastname  string
	age       int
	height    int
	weight    int
	hometown  string
	country   string
	retired   bool
}

//Sort all person by weight
type sortbyweight []person

func (sliceinfo sortbyweight) Len() int {
	//fmt.Println(len(slicelenght))
	return len(sliceinfo)
}
func (sliceinfo sortbyweight) Swap(i, j int)      { sliceinfo[i], sliceinfo[j] = sliceinfo[j], sliceinfo[i] }
func (sliceinfo sortbyweight) Less(i, j int) bool { return sliceinfo[i].weight < sliceinfo[j].weight }

//Sort all person by Height
type sortbyage []person

func (sliceinfo sortbyage) Len() int {
	return len(sliceinfo)
}
func (sliceinfo sortbyage) Swap(i, j int) {
	sliceinfo[i], sliceinfo[j] = sliceinfo[j], sliceinfo[i]
}
func (sliceinfo sortbyage) Less(i, j int) bool { return sliceinfo[i].age > sliceinfo[j].age }

func main() {

	person1 := person{
		firstname: "Rui",
		lastname:  "Hachimura",
		age:       22,
		height:    203,
		weight:    104,
		hometown:  "Toyama",
		country:   "Japan",
		retired:   false,
	}
	person2 := person{
		firstname: "Michael",
		lastname:  "Jordan",
		age:       57,
		height:    198,
		weight:    98,
		hometown:  "Brooklyn",
		country:   "USA",
		retired:   true,
	}

	person3 := person{
		firstname: "James",
		lastname:  "LeBron",
		age:       35,
		height:    203,
		weight:    113,
		hometown:  "Akron",
		country:   "USA",
		retired:   false,
	}

	person4 := person{
		firstname: "Tony",
		lastname:  "Parker",
		age:       38,
		height:    188,
		weight:    84,
		hometown:  "Villeurbane",
		country:   "France",
		retired:   true,
	}

	people := []person{person1, person2, person3, person4}
	fmt.Printf("Before sort by weight\n%v\n\n", people)
	sort.Sort(sortbyweight(people))
	fmt.Printf("After sort by weight\n%v\n\n", people)

	fmt.Printf("Before sort by age\n%v\n\n", people)
	sort.Sort(sortbyage(people))
	fmt.Printf("After sort by age\n%v\n\n", people)

	/*for _, value := range people {
		//fmt.Printf("Person ID #%d\n", indice)
		fmt.Printf("First Name:%s\nLast Name:%s\nAge:%d\nHeight:%d cm\nWeight:%d Kg\nHome Town:%s\nCountry:%s\nRetired:%t\n\n\n", value.firstname, value.lastname, value.age, value.height, value.weight, value.hometown, value.country, value.retired)

	}*/
}
