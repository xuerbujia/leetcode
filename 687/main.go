package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var hash = map[*TreeNode]int{}
	var depth func(root *TreeNode, target int) int
	var ans int
	depth = func(root *TreeNode, target int) int {
		if root == nil || root.Val != target {
			return 0
		}
		leftDepth := depth(root.Left, target)
		rightDepth := depth(root.Right, target)
		ans = max(ans, leftDepth+rightDepth)
		hash[root] = ans
		return max(leftDepth, rightDepth) + 1
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		depth(root, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
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
	longestUnivaluePath(&TreeNode{5, nil, &TreeNode{5, nil, &TreeNode{5, nil, nil}}})
}
