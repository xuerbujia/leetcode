package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := depth(root.Left)
		rightDepth := depth(root.Right)
		ans = max(ans, leftDepth+rightDepth)
		return max(leftDepth, rightDepth) + 1
	}
	depth(root)
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
