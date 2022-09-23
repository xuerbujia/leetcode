package main

type MyLinkedList struct {
	head *Node
	tail *Node
}
type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.Next
		if node == nil {
			return -1
		}
	}
	if node == nil {
		return -1
	}
	return node.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, l.head}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.tail == nil {
		l.AddAtHead(val)
		return
	}
	node := &Node{val, nil}
	l.tail.Next = node
	l.tail = node
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		l.AddAtHead(val)
		return
	}
	node := &Node{val, nil}
	h := l.head
	for i := 0; i < index-1; i++ {
		h = h.Next
		if h == nil {
			return
		}
	}
	if h == nil {
		return
	}
	node.Next = h.Next
	h.Next = node
	if node.Next == nil {
		l.tail = node
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		l.head = l.head.Next
		return
	}
	node := l.head
	for i := 0; i < index-1; i++ {
		node = node.Next
		if node == nil {
			return
		}
	}
	if node.Next == nil {
		return
	}
	node.Next = node.Next.Next
	if node.Next == nil {
		l.tail = node
	}
}
func main() {
	l := Constructor()
	l.AddAtIndex(1, 0)
	//l.DeleteAtIndex(1)
	//l.AddAtHead(1)
	//l.AddAtHead(5)
}
