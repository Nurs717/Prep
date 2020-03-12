package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 9}
	var target int = 8
	fmt.Println(twosums(nums, target))
}

func twosums(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
