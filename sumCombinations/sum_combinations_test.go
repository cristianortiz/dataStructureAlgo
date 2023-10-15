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
		combination := SumCombinations(nums, target)
		expected := [][]int{{1, 2, 3}, {1, 5}}
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})
	t.Run("find valid combinations", func(t *testing.T) {
		nums := []int{1, 2, 1, 1, 1, 1, 3, 1}
		target := 4
		combination := SumCombinations(nums, target)
		expected := [][]int{{1, 1, 1, 1}, {1, 1, 2}, {1, 3}}
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})
	//no solution
	t.Run("find valid combinations", func(t *testing.T) {
		nums := []int{1, 3, 434, 8}
		target := 5000
		combination := SumCombinations(nums, target)
		//empty slice of slices
		var expected [][]int
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})
}

func TestSumCombinations2(t *testing.T) {
	t.Run("return empty slice", func(t *testing.T) {
		nums := []int{0, 2}
		target := 1
		combination := SumCombinations(nums, target)
		//empty slice of slices
		var expected [][]int
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})

	t.Run("find valid combinations in sumCombinations2", func(t *testing.T) {
		nums := []int{2, 3, 6, 7}
		target := 7
		combination := SumCombinations2(nums, target)
		expected := [][]int{{2, 2, 3}, {7}}
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})

	t.Run("find valid combinations in sumCombinations2", func(t *testing.T) {
		nums := []int{2, 3, 5}
		target := 8
		combination := SumCombinations2(nums, target)
		expected := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
		if !reflect.DeepEqual(combination, expected) {
			t.Errorf("%v was expected, but received %v", expected, combination)
		}

	})

}
