package main

import (
	validparentheses "dataStructureAlgo/validParentheses"
)

func main() {
	// candidates := []int{2, 3, 6, 7}
	// target := 7
	// combination := sumcombinations.SumCombinations(candidates, target)
	// fmt.Println("Resultado final: ", combination)

	s := "()"
	valid := validparentheses.ValidParentheses(s)
	println("size input", len(s))

	println("response", valid)
}
