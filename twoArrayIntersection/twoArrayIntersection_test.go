package twoarrayintersection

import "testing"

func TestTwoArrayIntersection(t *testing.T) {
	if !arrayEquals(TwoArrayIntersection([]int{1, 2, 2, 1}, []int{2, 2}), []int{2}) {
		t.Error("solution is incorrect")
	}
	if !arrayEquals(TwoArrayIntersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{9, 4}) {
		t.Error("solution is incorrect")
	}
	if !arrayEquals(TwoArrayIntersection([]int{1, 2, 2, 3}, []int{2, 2, 3, 3}), []int{2, 3}) {
		t.Error("solution is incorrect")
	}

}

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
