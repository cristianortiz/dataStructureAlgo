package containsDuplicate

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
// Example:

// Input: nums = [1,2,3,1]
// Output: true

// using the map[][] aproach to solve the problem usgin only one for loop
// for trun the test for this specific function use the command
// go test -run /Map
func ContainDuplicateWithMap(nums []int) bool {
	numMap := make(map[int]bool)

	for _, num := range nums {

		if _, ok := numMap[num]; ok {
			return true
		}
		numMap[num] = true

	}

	return false
}
