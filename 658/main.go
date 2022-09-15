package main

import (
	"fmt"
)

func findClosestElements(arr []int, k int, x int) []int {
	var idx int
	idx = search(arr, x)
	var ans = make([]int, len(arr))
	var mid, left, right = idx, idx - 1, idx + 1
	if left < 0 {
		left = mid
		right = mid + 1
	} else if right == len(arr) {
		right = mid
		left = mid - 1
	} else {
		if abs(arr[left]-x) < abs(arr[right]-x) || (abs(arr[left]-x) == abs(arr[right]-x) && arr[left] < arr[right]) {
			left = mid - 1
			right = mid
		} else {
			left = mid
			right = mid + 1
		}
	}

	for right-left-1 < k {
		if right == len(arr) || (left >= 0 && (abs(arr[left]-x) < abs(arr[right]-x) || (abs(arr[left]-x) == abs(arr[right]-x) && arr[left] < arr[right]))) {
			ans[left] = arr[left]
			left--
		} else {
			ans[right] = arr[right]
			right++
		}
	}
	return ans[left+1 : right]
}
func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
func search(arr []int, target int) int {
	var left, right = 0, len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
func main() {
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}
