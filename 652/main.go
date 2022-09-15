package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var ans []*TreeNode
	var hash = map[string]int{}
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		s := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
		if v := hash[s]; v == 1 {
			ans = append(ans, root)
		} else {
			hash[s]++
		}
		return s
	}
	dfs(root)
	return ans
}
func main() {
	a := "bbb"
	var hash = map[int]*string{1: &a}

	c := "bbb"
	fmt.Println(*hash[1] == c)
}
