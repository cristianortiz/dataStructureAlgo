package selectionsort

func selection_sort(arr []int) []int {

	//external loop, in every passthrough the current index slice element will be the default lowest value,
	//in first passthrough arr[0] is the default lowest value
	for i := 0; i < len(arr); i++ {
		lowestNumberIndex := i
		//internal loop, In every pass-through, the current 'j' index is compared with the adjacent element of it
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowestNumberIndex] {
				//new lowest element founded
				lowestNumberIndex = j
			}

		}
		// If a new lowest element has been found in this passthrough, swap it with the previous lowest number.
		// If not, the current passthrough lowest value remains
		arr[i], arr[lowestNumberIndex] = arr[lowestNumberIndex], arr[i]

	}
	return arr
}
