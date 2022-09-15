package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var ans [][]int
	var flag = false
	var left, right = 0, intervals[0][1]
lab:
	for i := 0; i < len(intervals); i++ {
		if newInterval[0] >= left && newInterval[0] <= right {
			flag = true
			for j := i + 1; j <= len(intervals); j++ {
				if j == len(intervals) {
					ans = append(ans, []int{min(intervals[i][0], newInterval[0]), max(intervals[j-1][1], newInterval[1])})
					break lab
				}
				if intervals[j][0] > newInterval[1] {
					ans = append(ans, []int{intervals[i][0], newInterval[1]})
					ans = append(ans, intervals[j:]...)
					break lab
				}
				if intervals[j][0] == newInterval[1] {
					ans = append(ans, []int{intervals[i][0], intervals[j][1]})
					ans = append(ans, intervals[j+1:]...)
					break lab
				}
			}
		}
		ans = append(ans, intervals[i])
		if i < len(intervals)-1 {
			left = intervals[i+1][0]
			right = intervals[i+1][1]
		}
	}
	if !flag {
		if newInterval[1] < intervals[0][0] {

			ans = append([][]int{newInterval}, ans...)
		} else {
			ans = append(ans, newInterval)
		}
	}
	//	fmt.Println(ans)
	return ans
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
	fmt.Println(insert([][]int{{1, 5}}, []int{0, 0}))
}
