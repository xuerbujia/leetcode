package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var ans [][]int

	queue = append(queue, root)
	var flag = true
	for len(queue) > 0 {
		size := len(queue)
		var temp []int
		for i := 0; i < size; i++ {
			r := queue[0]
			queue = queue[1:]
			temp = append(temp, r.Val)
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
		if flag {
			ans = append(ans, temp)
		} else {
			reverse(temp)
			ans = append(ans, temp)
		}
		flag = !flag
	}
	return ans
}
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
func main() {
	nums := []int{1, 2}
	fmt.Printf("%p,%p", nums, nums[0])
	reverse(nums)
	fmt.Println(nums)
}
