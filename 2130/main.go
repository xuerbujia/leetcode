package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var reverse = new(ListNode)
	for slow != nil {
		temp := slow.Next
		slow.Next = reverse.Next
		reverse.Next = slow
		slow = temp
	}
	reverse = reverse.Next
	var ans int
	for reverse != nil {
		ans = max(ans, reverse.Val+head.Val)
		reverse = reverse.Next
		head = head.Next
	}
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

}
