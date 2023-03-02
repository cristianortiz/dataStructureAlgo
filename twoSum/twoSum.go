package twosum

// func main() {

// --- Problem Description--
//  Given an array of integers nums and an integer target, return indices of the two numbers
//  such that they add up to target. You may assume that each input would have exactly one solution,
// 	and you may not use the same element twice. You can return the answer in any order.
// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, the return slice must be [0, 1].

// using map[][] aproach to solve the problem in one for loop
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
