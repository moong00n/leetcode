package main

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test case 1",
			input:    "leet**cod*e",
			expected: "lecoe",
		},
		{
			name:     "test case 2",
			input:    "erase*****",
			expected: "",
		},
		{
			name:     "test case 3",
			input:    "helloo*world*****",
			expected: "hello",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solution(tc.input)
			if result != tc.expected {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}
