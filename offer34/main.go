package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var bfs func(root *TreeNode, target int)
	var temp []int
	var ans [][]int
	bfs = func(root *TreeNode, target int) {
		//fmt.Println(target)
		//target -= root.Val
		//if root.Left == nil && root.Right == nil {
		//	if target == 0 {
		//		ans = append(ans, append([]int{}, temp...))
		//
		//	}
		//
		//	return
		//}
		//
		//temp = append(temp, root.Val)
		//
		//bfs(root.Left, target)
		//bfs(root.Right, target)
		//
		//temp = temp[:len(temp)-1]
	}
	bfs(root, target)
	return ans
}
func main() {

}
