package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var left, right = -1, 0
	var hash = map[byte]bool{}
	var ans int
	for left < right && right < len(s) {
		for right < len(s) && !hash[s[right]] {
			hash[s[right]] = true
			right++
		}

		left++
		ans = max(ans, right-left)
		hash[s[left]] = false

	}
	fmt.Println(ans)
	return ans
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {
	lengthOfLongestSubstring("pwwkew")
}
