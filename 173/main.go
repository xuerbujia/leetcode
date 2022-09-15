package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	order []int
	idx   int
}

func Constructor(root *TreeNode) BSTIterator {
	var order []int
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		order = append(order, root.Val)
		helper(root.Right)
	}
	helper(root)
	//fmt.Println(order)
	return BSTIterator{order: order}
}

func (this *BSTIterator) Next() int {
	this.idx++
	return this.order[this.idx-1]
}

func (this *BSTIterator) HasNext() bool {
	return this.idx < len(this.order)
}
func main() {
	root := &TreeNode{7, &TreeNode{3, nil, nil}, &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	Constructor(root)
}
