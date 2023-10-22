package validparentheses

import "testing"

type testCase struct {
	input    string
	expected bool
}

func TestValidParentheses(t *testing.T) {

	trails := []testCase{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"{{[()]}}", true},
		{"[}", false},
		{"]", false},
		{"(", false},
	}

	for _, trail := range trails {
		t.Run(trail.input, func(t *testing.T) {
			result := ValidParentheses(trail.input)
			if result != trail.expected {
				t.Errorf(" ValidParentheses = %v; expected %v", trail.input, result)
			}

		})
	}
}
