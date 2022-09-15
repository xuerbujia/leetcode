package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var inOrder []*TreeNode
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		inOrder = append(inOrder, root)
		helper(root.Right)
	}
	helper(root)
	var first, fnext *TreeNode
	var snext *TreeNode
	for i := 0; i < len(inOrder)-1; i++ {
		if inOrder[i].Val > inOrder[i+1].Val {
			if first == nil {
				first = inOrder[i]
				fnext = inOrder[i+1]
			} else {
				snext = inOrder[i+1]
			}
		}
	}
	if snext == nil {
		first.Val, fnext.Val = fnext.Val, first.Val
	} else {
		first.Val, snext.Val = snext.Val, first.Val
	}
}
func main() {
	recoverTree(&TreeNode{1, &TreeNode{3, nil, &TreeNode{2, nil, nil}}, nil})

}
