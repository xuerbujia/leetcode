package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var ans = new(ListNode)
	ansnode := ans
	fast := head
lab:
	for head != nil {
		for i := 0; i < k; i++ {
			if fast == nil {
				break lab
			}
			fast = fast.Next
		}
		reverseList, tail := reverseK(head, k)
		ansnode.Next = reverseList
		ansnode = tail
		head = fast
	}
	ansnode.Next = head
	return ans.Next
}
func reverseK(head *ListNode, k int) (reverseList, tail *ListNode) {
	reverseList = new(ListNode)
	tail = head
	for i := 0; i < k; i++ {
		temp := head.Next
		head.Next = reverseList.Next
		reverseList.Next = head
		head = temp
	}
	return reverseList.Next, tail
}
func main() {

}
