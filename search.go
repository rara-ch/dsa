package main

import (
	"math"
)

func linearSearch(s []int, target int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return i
		}
	}
	return -1
}

func binarySearch(s []int, target int) int {
	i := 0
	j := len(s) - 1

	for i <= j {
		m := (i + j) / 2

		if s[m] == target {
			return m
		} else if target < s[m] {
			j = m - 1
		} else if target > s[m] {
			i = m + 1
		}
	}
	return -1
}

func recursiveBinarySearch(s []int, left int, right int, target int) int {
	m := (left + right) / 2

	if left > right {
		return -1
	} else if target == m {
		return m
	} else if target < m {
		return recursiveBinarySearch(s, left, m-1, target)
	} else {
		return recursiveBinarySearch(s, m+1, right, target)
	}
}

func firstOccuranceBinarySearch(s []int, target int) int {
	firstOccurance := -1
	left := 0
	right := len(s) - 1

	for left <= right {
		m := (left + right) / 2

		if s[m] == target {
			firstOccurance = m
			right = m - 1
		} else if s[m] < target {
			left = m + 1
		} else if s[m] > target {
			right = m - 1
		}
	}

	return firstOccurance
}

func lastOccuranceBinarySearch(s []int, target int) int {
	lastOccurance := -1
	left := 0
	right := len(s) - 1

	for left <= right {
		m := (left + right) / 2

		if s[m] == target {
			lastOccurance = m
			left = m + 1
		} else if s[m] < target {
			left = m + 1
		} else if s[m] > target {
			right = m - 1
		}
	}

	return lastOccurance
}

func lowerBoundBinarySearch(s []int, target int) int {
	bound := -1
	left := 0
	right := len(s) - 1

	for left <= right {
		m := (left + right) / 2

		if s[m] <= target {
			left = m + 1
		} else if s[m] > target {
			bound = m
			right = m - 1
		}
	}

	return bound
}

func upperBoundBinarySearch(s []int, target int) int {
	bound := -1
	left := 0
	right := len(s) - 1

	for left <= right {
		m := (left + right) / 2

		if s[m] >= target {
			right = m - 1
		} else if s[m] < target {
			bound = m
			left = m + 1
		}
	}

	return bound
}

// Tested on LeetCode (Problem 34)
func firstLastBinarySearch(s []int, target int) (int, int) {
	left := 0
	right := len(s) - 1
	pivot := -1

	for left <= right {
		m := (left + right) / 2

		if s[m] == target {
			pivot = m
			break
		} else if s[m] < target {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	if pivot == -1 {
		return -1, -1
	}

	left = 0
	right = pivot - 1
	firstOccurance := pivot

	for left <= right {
		m := (left + right) / 2

		if s[m] == target {
			firstOccurance = m
			right = m - 1
		} else if s[m] > target {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	left = pivot + 1
	right = len(s) - 1
	lastOccurance := pivot

	for left <= right {
		m := (left + right) / 2

		if s[m] == target {
			lastOccurance = m
			left = m + 1
		} else if s[m] > target {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return firstOccurance, lastOccurance
}

// Tested on LeetCode (Problem 162)
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2

		lessLeft := false
		lessRight := false

		if (m+1) > (len(nums)-1) || nums[m] > nums[m+1] {
			lessRight = true
		}

		if (m-1) < 0 || nums[m] > nums[m-1] {
			lessLeft = true
		}

		if lessLeft && lessRight {
			return m
		} else if lessLeft {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return 0
}

func twoCrystalBallProblem(heights []bool) int {
	i := 0
	n := len(heights)
	var j int
	for i < n {
		j = i
		if heights[i] {
			break
		}
		i += int(math.Sqrt(float64(n)))
	}

	for j < n {
		if heights[j] {
			return j
		}
		j++
	}

	return -1
}
