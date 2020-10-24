package main

import "testing"

func TestSum(t *testing.T) {
	//Create a structure to defien values used for the test
	type test struct {
		testdataslice  []int
		testdataresult int
	}
	//Create a slice of type test
	//Assign a set of test values
	testdata := []test{
		test{[]int{4, 4}, 8},
		test{[]int{2, 2}, 4},
		test{[]int{3, 3}, 6},
		test{[]int{23, 24}, 47},
	}

	for _, value := range testdata {
		result := sum(value.testdataslice...)
		if result != value.testdataresult {
			t.Error("Result expected is", value.testdataresult, "But we obtained", result)
		}
	}
}
