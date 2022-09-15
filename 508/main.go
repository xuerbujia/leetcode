package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var hash = map[int]int{}
	var max int
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := helper(root.Left)
		rightSum := helper(root.Right)
		sum := leftSum + rightSum + root.Val
		hash[sum]++
		if hash[sum] > max {
			max = hash[sum]
		}
		return sum
	}
	helper(root)
	var ans []int
	for k, v := range hash {
		if v == max {
			ans = append(ans, k)
		}
	}
	return ans
}
func main() {

}
