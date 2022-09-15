package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(node *ListNode, root *TreeNode) bool
	dfs = func(node *ListNode, root *TreeNode) bool {
		if node == nil {
			return true
		}
		if root == nil || node.Val != root.Val {
			return false
		}

		return dfs(node.Next, root.Left) || dfs(node.Next, root.Right)

	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
func main() {
	head := &ListNode{1, &ListNode{10, nil}}

	right := &TreeNode{1, &TreeNode{10, &TreeNode{9, nil, nil}, nil}, &TreeNode{1, nil, nil}}
	root := &TreeNode{1, nil, right}
	fmt.Println(isSubPath(head, root))
}
