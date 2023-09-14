package bubblesort

func bubble_sort(arr []int) []int {

	n := len(arr)
	swapped := true
	for swapped {
		//if a passtrough is completed without any elements swapped, it means the entire
		//slice is ordered correctly, and the for loops ends
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				//swap the elements
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}

		}

	}
	return arr
}
