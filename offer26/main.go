package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	root := B
	var bfs func(A, B *TreeNode, searchB bool) bool
	bfs = func(A, B *TreeNode, searchB bool) bool {
		if B == nil {
			return true
		}
		if searchB && B.Val != A.Val {
			return false
		}
		if B.Val == A.Val {
			if bfs(A.Left, B.Left, true) && bfs(A.Right, B.Right, true) {
				return true
			}
		}

		if A == nil {
			return false
		}

		if bfs(A.Left, root, false) || bfs(A.Right, root, false) {
			return true
		}
		return false
	}
	if bfs(A, B, false) {
		return true
	}
	return false
}
func main() {
	var a = new(TreeNode)
	a.Val = 3
	root := a
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root = root.Left
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	var b = new(TreeNode)
	b.Val = 4
	isSubStructure(a, b)
}
