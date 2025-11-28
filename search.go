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
