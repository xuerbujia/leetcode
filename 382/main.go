package main

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}
type Solution struct {
	List []*ListNode
}

func Constructor(head *ListNode) Solution {
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	return Solution{list}
}

func (this *Solution) GetRandom() int {
	return this.List[rand.Intn(len(this.List))].Val
}
func main() {

}
