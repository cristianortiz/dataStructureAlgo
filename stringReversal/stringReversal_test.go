package stringreversal

import (
	"reflect"
	"testing"
)

func TestReverseStringRecursive(t *testing.T) {

	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		//{[]byte{'y', 'o', 'u', ',', ' ', 's', 'o', 'n'}, []byte{'n', 'o', 's', ',', 'u', 'o', 'y'}},
		{[]byte{'a', 'b', 'c'}, []byte{'c', 'b', 'a'}},
		{[]byte{'1'}, []byte{'1'}},
		{[]byte{}, []byte{}},
	}

	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			ReverseStringRecursive(tc.input)
			result := tc.input
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
func TestReverseStringIterative(t *testing.T) {

	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'y', 'o', 'u', ',', ' ', 's', 'o', 'n'}, []byte{'n', 'o', 's', ' ', ',', 'u', 'o', 'y'}},
		{[]byte{'a', 'b', 'c'}, []byte{'c', 'b', 'a'}},
		{[]byte{'1'}, []byte{'1'}},
		{[]byte{}, []byte{}},
	}

	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			ReverseStringIterative(tc.input)
			result := tc.input
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
