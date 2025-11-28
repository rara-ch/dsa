package main

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
