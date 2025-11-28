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
