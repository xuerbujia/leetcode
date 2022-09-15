package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	var ans = new(ListNode)
	ans.Next = head
	node := ans
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
	return ans.Next
}
func main() {

}
