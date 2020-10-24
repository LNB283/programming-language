package main

import "testing"

func TestSum(t *testing.T) {
	testvalue := sum(1, 1)
	if testvalue != 2 {
		t.Error("Result expected :", 2, "but got:", testvalue)
	}
}
