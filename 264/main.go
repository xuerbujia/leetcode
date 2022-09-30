package main

import (
	"container/heap"
	"fmt"
)

type hp []int

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	p := *h
	return p[i] < p[j]
}

func (h *hp) Swap(i, j int) {
	p := *h
	p[i], p[j] = p[j], p[i]
}

func (h *hp) Push(x interface{}) {
	p := *h
	p = append(p, x.(int))
	*h = p
}

func (h *hp) Pop() interface{} {
	p := *h
	res := p[len(p)-1]
	*h = p[:len(p)-1]
	return res
}

var vec = []int{2, 3, 5}

func nthUglyNumber(n int) int {
	h := new(hp)
	heap.Init(h)
	var m = map[int]bool{}
	heap.Push(h, 1)
	var ans int
	for idx := 0; idx < n; idx++ {
		ans = heap.Pop(h).(int)
		for _, v := range vec {
			num := v * ans
			if !m[num] {
				m[num] = true
				heap.Push(h, num)
			}
		}
	}
	fmt.Println(ans)
	return ans
}
func main() {
	nthUglyNumber(3)
}
