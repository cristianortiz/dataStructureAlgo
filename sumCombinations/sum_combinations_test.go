package sumcombinations

import (
	"reflect"
	"testing"
)

// Test solutions for TwoSum problem using map[][] aproach, using three predifined
// test cases and their correct outputs
func TestSumCombinations(t *testing.T) {
	t.Run("find valid combinations", func(t *testing.T) {
		nums := []int{1, 5, 3, 2}
		target := 6
		combination := sumCombinations(nums, target)
		expected := [][]int{{1, 2, 3}, {1, 5}}
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})
	t.Run("find valid combinations", func(t *testing.T) {
		nums := []int{1, 2, 1, 1, 1, 1, 3, 1}
		target := 4
		combination := sumCombinations(nums, target)
		expected := [][]int{{1, 1, 1, 1}, {1, 1, 2}, {1, 3}}
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})
	//no solution
	t.Run("find valid combinations", func(t *testing.T) {
		nums := []int{1, 3, 434, 8}
		target := 5000
		combination := sumCombinations(nums, target)
		//empty slice of slices
		var expected [][]int
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})
}
