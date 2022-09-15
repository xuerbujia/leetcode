package main

import "fmt"

func reformat(s string) string {
	var word []byte
	var nums []byte
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			nums = append(nums, s[i])
		} else {
			word = append(word, s[i])
		}
	}
	if abs(len(nums)-len(word)) > 1 {
		return ""
	}
	if len(nums) > len(word) {
		return string(getSlice(nums, word))
	} else {
		return string(getSlice(word, nums))
	}
}
func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
func getSlice(s1, s2 []byte) []byte {
	var ans = make([]byte, len(s1)+len(s2))
	var l1, l2, idx int
	for idx < len(ans) {
		if l1 < len(s1) {
			ans[idx] = s1[l1]
			l1++
			idx++
		}
		if l2 < len(s2) {
			ans[idx] = s2[l2]
			l2++
			idx++
		}
	}
	return ans
}

func main() {
	fmt.Println(reformat("covid219"))
}
