package insertionsort

func insertion_sort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		temp_value := arr[i]
		j := i

		for j >= 0 && arr[j-1] > temp_value {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		arr[j] = temp_value

	}
	return arr
}
