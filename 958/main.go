package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if !left && right {
			return false
		}
		return true
	}
	return dfs(root)
}
func main() {
	var a = make([]int, 0, 10)
	test(a)
	a = append(a, 2)
	fmt.Println(a)
}
func test(nums []int) {
	b := 1
	nums = append(nums, b)
}
