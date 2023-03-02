package twosum

// func main() {

// 	//--- Problem Description--
// 	// Given an array of integers nums and an integer target, return indices of the two numbers
// 	// such that they add up to target. You may assume that each input would have exactly one solution,
// 	// and you may not use the same element twice. You can return the answer in any order.
// 	data := []int{2, 7, 11, 15}
// 	target := 9
// 	res := twoSumWithMap(data, target)
// 	fmt.Println(res)
// }

func TwoSumWithMap(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i

	}

	return nil
}
