package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func isNumber(s string) bool {
	s = strings.ToLower(s)
	var dot, e = 20, 20
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if dot != 20 {
				return false
			}
			dot = i
		}
		if s[i] == 'e' {
			if e != 20 {
				return false
			}
			e = i
		}
	}
	nums := strings.Split(s, ".")
	if len(nums) > 1 && nums[0] == "" && nums[1] == "" {
		return false
	}
	if len(nums[0]) == 1 && len(nums) > 1 && nums[1] == "" {
		if nums[0][0] < '0' || nums[0][0] > '9' {
			return false
		}
	}
	if !(dot == 20 || e == 20) {
		if e < dot {
			return false
		}
	}

	var haveE = false
	for i, v := range nums {
		n := strings.Split(v, "e")
		if len(n) == 2 {
			if haveE || n[0] == "" || n[1] == "" {
				return false
			}
			if len(n[0]) == 1 || len(n[1]) == 1 {
				if n[0][0] < '0' || n[0][0] > '9' || n[1][0] < '0' || n[1][0] > '9' {
					return false
				}
			}
			haveE = true
			//if i == 0 {
			//	ok, _ := regexp.MatchString(`^[+-]{0,1}\d{0,}$`, n[0])
			//	if !ok {
			//		return false
			//	}
			//} else {
			//	ok, _ := regexp.MatchString(`^\d{0,}$`, n[0])
			//	if !ok {
			//		return false
			//	}
			//}
			//ok, _ := regexp.MatchString(`^[+-]{0,1}\d{0,}$`, n[1])
			//if !ok {
			//	return false
			//}
			//continue
		}
		for j, v2 := range n {
			var pattern string
			if i == 0 && j == 0 {
				pattern = `^[+-]{0,1}\d{0,}$`
			} else {
				pattern = `^\d{0,}$`
			}
			ok, _ := regexp.MatchString(pattern, v2)
			if !ok {
				return false
			}
		}

	}
	return true
}

func main() {

	//	fmt.Println(isNumber("+E3"))
	usercardsSlice := []int{14, 1, 13, 2, 13, 3, 14, 4, 5, 15}

	var left, right = 0, len(usercardsSlice) - 1
	for left < right {
		for usercardsSlice[left] == 13 || usercardsSlice[left] == 14 {
			usercardsSlice[left], usercardsSlice[right] = usercardsSlice[right], usercardsSlice[left]
			right--

		}
		left++
	}
	count := len(usercardsSlice) - right
	fmt.Println(usercardsSlice, count)
	//sort.Ints(usercardsSlice[:len(usercardsSlice)-count])
	sort.Slice(usercardsSlice[:len(usercardsSlice)-count], func(i, j int) bool {
		return usercardsSlice[i] > usercardsSlice[j]
	})
	fmt.Println(usercardsSlice)
}
