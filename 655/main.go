package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getPow(j int) int {
	var ans = 1
	for i := 0; i < j; i++ {
		ans *= 2
	}
	return ans
}
func printTree(root *TreeNode) [][]string {
	var helper func(deep int, root *TreeNode)
	var d int
	helper = func(deep int, root *TreeNode) {
		if root == nil {
			return
		}
		if deep > d {
			d = deep
		}
		helper(deep+1, root.Left)
		helper(deep+1, root.Right)
	}
	helper(0, root)

	m := d + 1
	n := getPow(m) - 1
	var ans = make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}

	var row, col = 0, (n - 1) / 2
	ans[row][col] = strconv.Itoa(root.Val)
	var dfs func(row, col int, root *TreeNode)
	dfs = func(row, col int, root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			ans[row+1][col-getPow(d-1-row)] = strconv.Itoa(root.Left.Val)
			dfs(row+1, col-getPow(d-1-row), root.Left)
		}
		if root.Right != nil {
			ans[row+1][col+getPow(d-1-row)] = strconv.Itoa(root.Right.Val)
			dfs(row+1, col+getPow(d-1-row), root.Right)
		}
	}
	dfs(row, col, root)
	return ans
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	printTree(root)
}
