package sumcombinations

import (
	"fmt"
	"sort"
)

// Given an array of distinct integers candidates and a target integer target, return a list of all unique
// combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.
// The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
// frequency  of at least one of the chosen numbers is different. The test cases are generated such that the
// number of unique combinations that sum up to target is less than 150 combinations for the given input.
// Example 1:

// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.
// Example 2:

// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]
// Example 3:

// Input: candidates = [2], target = 1
// Output: []

func SumCombinations2(candidates []int, target int) [][]int {
	//valid combinations
	var result [][]int
	//state of combinations
	var combinations []int

	//ordering candidates slice
	sort.Ints(candidates)
	call := 1

	//recursive funtions definition
	var backtrack func(remain, start int)
	backtrack = func(remain, start int) {
		if remain == 0 {
			result = append(result, append([]int{}, combinations...))
			fmt.Printf("combinacion VALIDA: %v \n", result)

			return
		}
		for i := start; i < len(candidates); i++ {

			if candidates[i] > remain {
				fmt.Printf("candidate[%v]= %v ,es mayor que remain= %v \n", i, candidates[i], remain)
				continue
			}

			if remain-candidates[i] < 0 {
				break
			}
			combinations = append(combinations, candidates[i])
			//call the backtrack funcion recursively with updated paremeters
			call++
			fmt.Printf("%v- backtrack(remain=%v, start=%v, combination=%v)\n", call, remain-candidates[i], start, combinations)

			backtrack(remain-candidates[i], i)
			combinations = combinations[:len(combinations)-1]
			fmt.Printf("quitar ultimo, nuevo combination: %v \n", combinations)
		}

	}
	backtrack(target, 0)
	return result
}

func SumCombinations3(candidates []int, target int) [][]int {
	var result [][]int
	var current []int
	backtrack(&result, &current, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, current *[]int, candidates []int, target, start int) {
	if target == 0 {
		// If we have found a combination that sums up to the target, add it to the result.
		*result = append(*result, append([]int{}, (*current)...))
		fmt.Printf("combinacion VALIDA: %v \n", result)

		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			// If the current candidate is larger than the remaining target, skip it.
			continue
		}

		// Include the current candidate in the combination.
		*current = append(*current, candidates[i])

		// Recursively explore combinations starting from the current candidate.
		fmt.Printf("backtrack(result=%v, current=%v, candidates=%v, target=%v,start=%v)\n", result, current, candidates, target-candidates[i], i)

		backtrack(result, current, candidates, target-candidates[i], i)

		// Backtrack by removing the current candidate.
		*current = (*current)[:len(*current)-1]
		fmt.Printf("   quitar ultimo, nuevo combination: %v \n", current)

	}
}
