package main

import "fmt"

func main() {
	index := firstOccuranceBinarySearch([]int{1, 2, 2, 2, 3, 4}, 5)

	fmt.Println(index)
}
