package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {

	arr := []int{11, 21, 33, 41, 50, 69, 72, 89, 93, 1004, 1134, 1266, 13788}
	t.Run("Find the middle element position", func(t *testing.T) {
		result := binarySearch(arr, 72)
		if result != 6 {
			t.Errorf(" 6 was expected but,  %d returned instead", result)
		}
	})
	t.Run("Find the first element position", func(t *testing.T) {
		result := binarySearch(arr, 11)
		if result != 0 {
			t.Errorf("position 0 was expected , but %d returned instead", result)
		}
	})

	t.Run("Find the position of the last element", func(t *testing.T) {
		result := binarySearch(arr, 13788)
		if result != 12 {
			t.Errorf("position 12 was expected, but %d returned instead", result)
		}
	})

	t.Run("return -1 if the element does not belongs to the list", func(t *testing.T) {
		result := binarySearch(arr, 2342)
		if result != -1 {
			t.Errorf("value -1 as expected,  but %d returned instead", result)
		}
	})

}
