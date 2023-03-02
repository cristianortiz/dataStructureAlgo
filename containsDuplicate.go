package main

import "fmt"

func containDuplicate() {

	//--- Problem Description--
	// Given an array of integers nums and an integer target, return indices of the two numbers
	// such that they add up to target. You may assume that each input would have exactly one solution,
	// and you may not use the same element twice. You can return the answer in any order.
	data := []int{2, 7, 11, 15}
	res := containDuplicateWithMap(data)
	fmt.Println(res)
}

func containDuplicateWithMap(nums []int) bool {
	numMap := make(map[int]bool)

	for _, num := range nums {

		if _, ok := numMap[num]; ok {
			return true
		}
		numMap[num] = true

	}

	return false
}
