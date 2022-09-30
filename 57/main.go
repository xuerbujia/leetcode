package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	var stack [][]int
	for _, v := range intervals {
		if len(stack) > 0 && v[0] < stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			v = []int{top[0], max(top[1], v[1])}
		}
		fmt.Println(v)
		stack = append(stack, v)
	}

	return stack
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
func main() {
	insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
}
