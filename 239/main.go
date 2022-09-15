package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {

	left, right := 0, k+1
	var queue []int
	var queueHelper = func(i int) {
		if len(queue) == 0 {

		} else if nums[queue[0]] <= nums[i] {
			queue = []int{}
		} else if nums[queue[len(queue)-1]] < nums[i] {
			for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
				queue = queue[:len(queue)-1]
			}
		}
		queue = append(queue, i)
	}
	for i := 0; i < k; i++ {
		queueHelper(i)
	}

	var ans []int
	for right <= len(nums) {
		ans = append(ans, nums[queue[0]])
		if left == queue[0] {
			queue = queue[1:]
		}
		queueHelper(right - 1)
		left++
		right++

	}

	return append(ans, nums[queue[0]])
}
func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
