package main

import "fmt"

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
	full  bool
}

func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{
		queue: make([]int, k),
		head:  0,
		tail:  0,
		full:  false,
	}
	return q
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsEmpty() {
		q.tail = 0
		q.head = 0
	}
	if q.IsFull() {
		return false
	} else {
		q.queue[q.tail] = value
		q.tail++
		if q.tail == len(q.queue) {
			q.tail = 0
			//q.full=true
		}
		if q.tail == q.head {
			q.full = true
		}
		return true
	}
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	} else {
		q.queue[q.head] = -1
		q.head++
		q.full = false
		if q.head == len(q.queue) {
			q.head = 0
		}
		return true
	}
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return 0
	}
	return q.queue[q.head]
}

func (q *MyCircularQueue) Rear() int {
	var tail int
	if q.tail == 0 {
		tail = len(q.queue)
	} else {
		tail = q.tail
	}
	return q.queue[tail-1]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.head == q.tail && !q.IsFull()
}

func (q *MyCircularQueue) IsFull() bool {
	return q.full
}

func main() {
	q := Constructor(6)
	q.EnQueue(6)
	q.Rear()
	q.Rear()
	q.DeQueue()
	q.EnQueue(5)
	q.Rear()
	q.DeQueue()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.head, q.tail, q.queue)
	fmt.Println(q.Front())
}
