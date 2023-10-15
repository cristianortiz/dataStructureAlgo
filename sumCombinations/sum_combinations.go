package sumcombinations

// Create a function that finds all combinations of numbers from a list that sum up to the target value.
// - The function will receive a list of positive integers and a target value.
// - To obtain the combinations, each element of the list can only be used once (but there may be repeated elements in it).
// - Example: List = [1, 5, 3, 2], Target = 6
// - Solutions: [1, 5] and [1, 3, 2] (both combinations sum up to 6)
// -(If there are no combinations, return an empty list)

import (
	"fmt"
	"sort"
)

func SumCombinations(candidates []int, target int) [][]int {
	//final list of valid combinations
	var result [][]int
	//temporal list of combinations
	var combination []int
	//sort candidates to avoid repeated combinations in final result
	sort.Ints(candidates)

	call := 1

	//recursive function to find valid sums combinations
	//start represent  the index to init the valid combinations search
	//remain is the current obejctive value
	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		//solution finded: if remaining sum is 0 add the current combinatio n to the result slice
		if remain == 0 {
			result = append(result, append([]int{}, combination...))
			fmt.Printf("Valid Combination: %v \n", result)

			return
		}
		for i := start; i < len(candidates); i++ {
			if start == len(candidates) {
				break
			}
			//skip duplicates to avoid
			if i > start && candidates[i] == candidates[i-1] {
				fmt.Printf("candidates[%v]= %v equal to candidates[i-1]\n", candidates[i], candidates[i-1])

				continue
			}
			if remain-candidates[i] < 0 {
				break
			}
			//add the current candidate to the combination
			combination = append(combination, candidates[i])
			//call the backtrack funcion recursively with updated paremeters
			call++

			fmt.Printf("%v- backtrack(start=%v, remain=%v, combination=%v)\n", call, i+1, remain-candidates[i], combination)
			backtrack(i+1, remain-candidates[i])
			//remove the last added candidate
			combination = combination[:len(combination)-1]
			fmt.Printf("remove last candidate, new combination: %v \n", combination)

		}
	}
	fmt.Printf("1er call backtrack(start=%v, remain=%v, combination=%v)\n", 0, target, combination)

	backtrack(0, target)
	return result
}
