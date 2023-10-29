package quicksort

func quickSort(arr []int) []int {
	// recursion base case
	if len(arr) <= 1 {
		return arr
	}

	//define a pivot, first array element in this case
	pivot := arr[0]
	//slices to keep pointers in a subarray partition
	var left, right []int

	for _, item := range arr[1:] {
		if item <= pivot {
			left = append(left, item)
		} else {
			right = append(right, item)
		}

	}
	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
