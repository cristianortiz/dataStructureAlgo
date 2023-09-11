package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	t.Run("Find the middle element position", func(t *testing.T) {
		result := binarySearch(arr, 7)
		if result != 6 {
			t.Errorf("position 6 was expected but,  %d returned instead", result)
		}
	})
	t.Run("Find the first element position", func(t *testing.T) {
		result := binarySearch(arr, 1)
		if result != 0 {
			t.Errorf("position 0 was expected , but %d returned instead", result)
		}
	})

	t.Run("Find the position of the last element", func(t *testing.T) {
		result := binarySearch(arr, 13)
		if result != 12 {
			t.Errorf("Se esperaba 12, pero se obtuvo %d", result)
		}
	})

	t.Run("return -1 if the element does not belongs to the list", func(t *testing.T) {
		result := binarySearch(arr, 2342)
		if result != -1 {
			t.Errorf("Se esperaba -1, pero se obtuvo %d", result)
		}
	})

}
