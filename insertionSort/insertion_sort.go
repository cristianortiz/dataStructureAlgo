package insertionsort

func insertion_sort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		temp_value := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > temp_value {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp_value

	}
	return arr
}
