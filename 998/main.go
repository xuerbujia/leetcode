package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var helper func(root *TreeNode)
	var Inorder []int
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		Inorder = append(Inorder, root.Val)
		helper(root.Right)
	}
	helper(root)
	var nums = append(Inorder, val)
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left < 0 || right >= len(nums) || left > right {
			return nil
		}
		var max, idx int
		for i := left; i <= right; i++ {
			if nums[i] > max {
				max = nums[i]
				idx = i
			}
		}
		r := &TreeNode{max, nil, nil}
		r.Left = buildTree(left, idx-1)
		r.Right = buildTree(idx+1, right)
		return r
	}
	return buildTree(0, len(nums)-1)
}
func main() {

}
