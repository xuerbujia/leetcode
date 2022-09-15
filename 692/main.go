package main

import (
	"container/heap"
)

type countAndWord struct {
	count int
	word  string
}
type smallHeap []countAndWord

func (h smallHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	}
	return h[i].word > h[j].word
}
func (h smallHeap) Len() int {
	return len(h)
}
func (h smallHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *smallHeap) Pop() interface{} {
	old := *h
	res := old[h.Len()-1]
	*h = old[:h.Len()-1]
	return res
}
func (h *smallHeap) Push(n interface{}) {
	*h = append(*h, n.(countAndWord))
}

func topKFrequent(words []string, k int) []string {
	var hash = map[string]int{}
	for _, v := range words {
		hash[v]++
	}

	h := &smallHeap{}
	heap.Init(h)
	for key, v := range hash {
		heap.Push(h, countAndWord{v, key})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var ans = make([]string, k)
	for i := h.Len() - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(countAndWord).word
	}
	return ans
}
func main() {
	words := []string{"the", "sunny", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	topKFrequent(words, 2)
}
