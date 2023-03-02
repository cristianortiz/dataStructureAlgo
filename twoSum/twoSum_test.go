package twosum

import "testing"

// arrayEquals compare two slices looping and comparing element by element trougth them, returns true
// if slices are equals and false if they did'nt
func arrayEquals(a, b []int) bool {
	//returns false if the the two slice are not of the same size
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Test solutions for TwoSum problem using map[][] aproach, using three predifined
// test cases and their correct outputs
func TestTwoSumWithMap(t *testing.T) {

	if !arrayEquals(TwoSumWithMap([]int{2, 7, 11, 15}, 9), []int{0, 1}) {
		t.Error("Solution is incorrect")

	}
	if !arrayEquals(TwoSumWithMap([]int{3, 2, 4}, 6), []int{1, 2}) {
		t.Error("Solution is incorrect")

	}
	if !arrayEquals(TwoSumWithMap([]int{3, 3}, 6), []int{0, 1}) {
		t.Error("Solution is incorrect")

	}
}
