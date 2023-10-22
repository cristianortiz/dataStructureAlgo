package validparentheses

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid. An input string is valid if:
// - Open brackets must be closed by the same type of brackets.
// - Open brackets must be closed in the correct order.
// - Every close bracket has a corresponding open bracket of the same type.
// Example 1:
// Input: s = "()"
// Output: true
//
// Example 2:
// Input: s = "()[]{}"
// Output: true
//
// Example 3:
// Input: s = "(]"
// Output: false
// --------------------------------
func ValidParentheses(s string) bool {
	charStack := make([]rune, 0)
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			charStack = append(charStack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(charStack) == 0 || charStack[len(charStack)-1] != bracketPairs[char] {
				return false
			}
			charStack = charStack[:len(charStack)-1]
		}
	}

	return len(charStack) == 0

}
