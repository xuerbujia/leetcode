package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	s := strconv.Itoa(num)
	n := []byte(s)
	for i := 0; i < len(n); i++ {
		if n[i] == '6' {
			n[i] = '9'
			break
		}
	}
	ans, _ := strconv.Atoi(string(n))
	return ans
}
func simplifyPath(path string) string {
	var stack []string
	var temp string
	for i := 0; i <= len(path); i++ {
		if i == len(path) || path[i] == '/' {
			if temp == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if temp == "." || temp == "" {

			} else {
				stack = append(stack, temp)
			}
			temp = ""
			continue
		}
		temp += string(path[i])

	}
	var ans = ""
	for _, v := range stack {
		ans += v + "/"
	}
	if ans != "" {
		ans = ans[:len(ans)-1]
	}

	return "/" + ans
}
func subarraySum(nums []int, k int) int {
	var prefix = make([]int, len(nums)+1)
	var ans int
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	var hash = map[int]int{}
	for i := 0; i < len(prefix); i++ {
		if v, ok := hash[prefix[i]-k]; ok {
			ans += v
		}
		hash[prefix[i]]++

	}

	//fmt.Println(prefix)
	return ans
}
func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
