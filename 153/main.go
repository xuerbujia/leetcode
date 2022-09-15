package main

import "fmt"

func findMin(nums []int) int {
	var left, right = 0, len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if nums[right] > nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
func main() {
	fmt.Println(findMin([]int{2, 1}))
}
