package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var path func(root *TreeNode) int
	var ans = math.MinInt
	path = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, path(root.Left))
		right := max(0, path(root.Right))
		lmr := root.Val + left + right
		ret := root.Val + max(left, right)
		ans = max(ans, max(lmr, ret))
		return ret
	}
	path(root)
	return ans
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {

}
