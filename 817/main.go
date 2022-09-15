package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	var hash = map[int]bool{}
	for _, v := range nums {
		hash[v] = true
	}
	var ans int
	var flag bool
	for node := head; node != nil; node = node.Next {
		if !hash[node.Val] {
			flag = true
		}
		if hash[node.Val] && !flag {
			ans++
			flag = !flag
		}
	}
	return ans
}
func main() {

}
