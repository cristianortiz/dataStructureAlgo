package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	t.Run("Order a list", func(t *testing.T) {
		unsortedSlice := []int{64, 34, 25, 12, 22, 11, 90}
		sorted := insertion_sort(unsortedSlice)
		expected := []int{11, 12, 22, 25, 34, 64, 90}
		if !reflect.DeepEqual(sorted, expected) {
			t.Errorf("%v was expected, but received %v", expected, sorted)
		}

	})

	t.Run("skip an ordered slice", func(t *testing.T) {
		sortedArray := []int{11, 12, 22, 25, 34, 64, 90}
		newArray := insertion_sort(sortedArray)
		if !reflect.DeepEqual(newArray, sortedArray) {
			t.Errorf("the slice has been altered: %v", newArray)
		}
	})

	t.Run("handle an empty slice", func(t *testing.T) {
		emptyArray := []int{}
		newArray := insertion_sort(emptyArray)
		if len(newArray) != 0 {
			t.Errorf("the slice has been altered: %v", newArray)
		}
	})
}
