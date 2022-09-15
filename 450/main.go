package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var helper func(root *TreeNode)
	var inOrder []int
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		inOrder = append(inOrder, root.Val)
		helper(root.Right)
	}
	helper(root)
	s := search(inOrder, key)
	if s == -1 {
		return root
	}
	nums := append(inOrder[:s], inOrder[s+1:]...)
	var buildTree func(nums []int, left, right int) *TreeNode
	buildTree = func(nums []int, left, right int) *TreeNode {
		if left < 0 || right >= len(nums) || left > right {
			return nil
		}
		mid := (left + right) / 2
		r := &TreeNode{nums[mid], nil, nil}
		r.Left = buildTree(nums, left, mid-1)
		r.Right = buildTree(nums, mid+1, right)
		return r
	}
	return buildTree(nums, 0, len(nums)-1)
}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	var n = []int{2, 3, 4, 5, 6, 7}
	i := search(n, 3)
	fmt.Println(i)
}
