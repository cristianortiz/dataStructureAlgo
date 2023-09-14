package binarysearch

func binarySearch(arr []int, target int) int {

	//defines array upper and lower positions bounds of a sorted list like this arr={1,2,3,4,5,6,7,8,9}
	left := 0
	right := len(arr) - 1

	//main loop untul bounds cross each other
	for left <= right {
		//Let's suppose I'm looking for the element = 8
		//middle point position between bounds (first loop) = 0 +(8-0)/2 = 4
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}
		//fist loop: arr[4] = 5 < 8 ?
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
