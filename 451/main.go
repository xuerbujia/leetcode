package main

import (
	"container/heap"
	"strings"
)

type peer struct {
	count int
	b     byte
}
type hp []peer

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i int, j int) bool {
	hp := *h
	return hp[i].count > hp[j].count
}
func (h *hp) Swap(i int, j int) {
	hp := *h
	hp[i], hp[j] = hp[j], hp[i]
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(peer))
}
func (h *hp) Pop() interface{} {
	old := *h
	res := old[old.Len()-1]
	*h = old[:old.Len()-1]
	return res
}

func frequencySort(s string) string {
	var count = map[byte]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	h := &hp{}
	heap.Init(h)
	for k, v := range count {
		heap.Push(h, peer{v, k})
	}
	var ans []byte
	for h.Len() > 0 {
		pop := heap.Pop(h)
		p := pop.(peer)
		repeat := strings.Repeat(string(p.b), p.count)
		ans = append(ans, repeat...)
	}
	return string(ans)
}
func main() {

}
