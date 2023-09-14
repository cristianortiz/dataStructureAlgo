package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	t.Run("Sort the list in ascending order", func(t *testing.T) {
		unsorted := []int{300, 250, 100, 95, 78, 50, 23, 9, 4}
		sorted := bubble_sort(unsorted)
		expected := []int{4, 9, 23, 50, 78, 95, 100, 250, 300}
		if !reflect.DeepEqual(sorted, expected) {
			t.Errorf("Error, %v was expected, %v received instead", expected, sorted)
		}

	})
}
