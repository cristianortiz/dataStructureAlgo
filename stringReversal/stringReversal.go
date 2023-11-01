package stringreversal

// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]

// this version is RECURSIVE and recieve a slice of bytes of chars as parameter
func ReverseStringRecursive(s []byte) {
	recursiveHelper(s, 0, len(s)-1)

}

func recursiveHelper(s []byte, left, right int) {
	// Caso base: cuando left >= right, no hay más elementos para intercambiar
	if left >= right {
		return
	}
	// Intercambia el elemento en left y right
	s[left], s[right] = s[right], s[left]

	// Llama recursivamente con los índices ajustados
	recursiveHelper(s, left+1, right-1)
}

// this version is iterative bus also works con []bytes as a parameter
func ReverseStringIterative(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
