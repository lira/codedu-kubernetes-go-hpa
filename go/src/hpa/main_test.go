package main

import "testing"

func TestSqrt(t *testing.T) {
	result := sqrt(64)
	var response float64 = 8

	if result != response {
		t.Errorf("Invalid result! :( return %f, wanted %f", result, response)
	}
}