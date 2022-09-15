package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val < low {
			return dfs(root.Left)
		}
		if root.Val > high {
			return dfs(root.Right)
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)
		return root
	}
	return dfs(root)
}
func main() {

}
