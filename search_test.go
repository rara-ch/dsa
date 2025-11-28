package main

import "testing"

func TestBinarySearch(t *testing.T) {
	basicArray := []int{0, 1, 2, 3, 4, 5, 6}

	type inputs struct {
		array  []int
		target int
	}

	cases := []struct {
		inputs   inputs
		expected int
	}{
		{
			inputs: inputs{
				array:  basicArray,
				target: 0,
			},
			expected: 0,
		},
		{
			inputs: inputs{
				array:  basicArray,
				target: 6,
			},
			expected: 6,
		},
		{
			inputs: inputs{
				array:  basicArray,
				target: 3,
			},
			expected: 3,
		},
		{
			inputs: inputs{
				array:  basicArray,
				target: 10,
			},
			expected: -1,
		},
		{
			inputs: inputs{
				array:  []int{},
				target: 1,
			},
			expected: -1,
		},
	}

	for _, c := range cases {
		got := binarySearch(c.inputs.array, c.inputs.target)
		if got != c.expected {
			t.Errorf("results do not match: got %v, expected %v", got, c.expected)
		}
	}
}
