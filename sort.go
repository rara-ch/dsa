package main

func swap(nums []int, i, j int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	return nums
}

func bubbleSort(nums []int) []int {
	for i := len(nums); i > 0; i-- {
		hasSwap := false
		for j := 0; j < i-1; j++ {
			if nums[j] > nums[j+1] {
				hasSwap = true
				swap(nums, j, j+1)
			}
		}
		if hasSwap {
			break
		}
	}
	return nums
}
