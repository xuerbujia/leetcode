package main

import (
	"container/heap"
	"fmt"
)

func getKthMagicNumber(k int) int {
	var h hp
	heap.Init(&h)
	heap.Push(&h, 1)
	var m = map[int]bool{}
	var idx int
	var ans int
	for idx < k {
		r := heap.Pop(&h).(int)
		idx++
		ans = r
		for i := 0; i < 3; i++ {
			var n int
			switch i {
			case 0:
				n = r * 3
			case 1:
				n = r * 5
			case 2:
				n = r * 7
			}
			if !m[n] {
				heap.Push(&h, n)

				m[n] = true
			}
		}

	}
	fmt.Println(ans)
	return ans
}

type hp []int

func (h hp) Len() int {
	return len(h)

}

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	p := h
	*p = append(*p, x.(int))
}

func (h *hp) Pop() interface{} {
	old := *h
	res := old[len(old)-1]

	*h = old[:len(old)-1]
	return res
}

func main() {
	getKthMagicNumber(7)
}
