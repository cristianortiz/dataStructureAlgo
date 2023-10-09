package sumcombinations

// Create a function that finds all combinations of numbers from a list that sum up to the target value.
// - The function will receive a list of positive integers and a target value.
// - To obtain the combinations, each element of the list can only be used once (but there may be repeated elements in it).
// - Example: List = [1, 5, 3, 2], Target = 6
// - Solutions: [1, 5] and [1, 3, 2] (both combinations sum up to 6)
// -(If there are no combinations, return an empty list)

import "sort"

func sumCombinations(candidates []int, target int) [][]int {
	//final list of valid combinations
	var result [][]int
	//temporal list of combinations
	var combination []int
	//sort candidates to avoid repeated combinations in final result
	sort.Ints(candidates)

	//recursive function to find valid sums combinations
	//start represent  the index to init the valid combinations search
	//remain is the current obejctive value
	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {

		//solution finded: if remaining sum is 0 add the current combination to the result slice
		if remain == 0 {
			result = append(result, append([]int{}, combination...))
			return
		}

		for i := start; i < len(candidates); i++ {

			//if the remaing sum is negative, there is no solution
			if target < 0 || start == len(candidates) {
				break
			}

			//skip duplicates to avoid
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			//add the current candidate to the combination
			combination = append(combination, candidates[i])
			//call the backtrack funcion recursively with updated paremeters
			backtrack(i+1, remain-candidates[i])
			//remove the last added candidate
			combination = combination[:len(combination)-1]
		}
	}

	//first call of recursive function
	backtrack(0, target)
	return result
}
