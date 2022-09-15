package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var helper func(nums []int, left, right int) *TreeNode
	helper = func(nums []int, left, right int) *TreeNode {
		if left > right || left < 0 || right >= len(nums) {
			return nil
		}
		if left == 4 && right == 4 {
			fmt.Println()
		}
		var max, maxidx = -1, 0

		for i := left; i <= right; i++ {
			if nums[i] > max {
				max = nums[i]
				maxidx = i
			}
		}
		root := &TreeNode{max, nil, nil}
		root.Left = helper(nums, left, maxidx-1)
		root.Right = helper(nums, maxidx+1, right)
		return root
	}
	return helper(nums, 0, len(nums)-1)

}
func main() {
	r := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(r)
}
