package main

import "testing"

type inputs struct {
	array  []int
	target int
}

func TestUniqueSearch(t *testing.T) {
	basicArray := []int{0, 1, 2, 3, 4, 5, 6}

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
		result := linearSearch(c.inputs.array, c.expected)
		if result != c.expected {
			t.Errorf("linear search results do not match: got %v, expected %v", result, c.expected)
		}

		result = binarySearch(c.inputs.array, c.inputs.target)
		if result != c.expected {
			t.Errorf("iterative binary search results do not match: got %v, expected %v", result, c.expected)
		}

		result = recursiveBinarySearch(c.inputs.array, 0, len(c.inputs.array)-1, c.inputs.target)
		if result != c.expected {
			t.Errorf("recursive binary search results do not match: got %v, expected %v", result, c.expected)
		}
	}
}

func TestFirstOccuranceBinarySearch(t *testing.T) {
	cases := []struct {
		inputs   inputs
		expected int
	}{
		{
			inputs: inputs{
				array:  []int{},
				target: 1,
			},
			expected: -1,
		},
		{
			inputs: inputs{
				array:  []int{1, 2, 2, 2, 2, 2, 3},
				target: 2,
			},
			expected: 1,
		},
		{
			inputs: inputs{
				array:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				target: 0,
			},
			expected: 0,
		},
	}

	for _, c := range cases {
		result := firstOccuranceBinarySearch(c.inputs.array, c.inputs.target)
		if result != c.expected {
			t.Errorf("first occurance binary search error: got %v, expected %v", result, c.expected)
		}
	}

}

func TestLastOccuranceBinarySearch(t *testing.T) {
	cases := []struct {
		inputs   inputs
		expected int
	}{
		{
			inputs: inputs{
				array:  []int{},
				target: 1,
			},
			expected: -1,
		},
		{
			inputs: inputs{
				array:  []int{1, 2, 2, 2, 2, 2, 3},
				target: 2,
			},
			expected: 5,
		},
		{
			inputs: inputs{
				array:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				target: 0,
			},
			expected: 12,
		},
	}

	for _, c := range cases {
		result := lastOccuranceBinarySearch(c.inputs.array, c.inputs.target)
		if result != c.expected {
			t.Errorf("last occurance binary search error: got %v, expected %v", result, c.expected)
		}
	}

}
