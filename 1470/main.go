package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var left, right = 1, n
	for i := 1; i < 2*n-1; i++ {
		if i%2 == 1 {
			nums[i] = nums[right]&0x3FF<<10 | nums[i]
			right++
		} else {
			nums[i] = nums[left]&0x3FF<<10 | nums[i]
			left++
		}
	}
	for i := 1; i < 2*n-1; i++ {
		nums[i] >>= 10
	}
	return nums
}
func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
