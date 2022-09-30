package main

import (
	"strings"
)

/*
简单模拟
先用一个map将数字和罗马字符映射建立
我们可以通过对数字可以分成几个1000、100、10、1来将数字分为 A千B百C十D 这样的形式，
得到这样的形式以后就可以对用A、B、C、D和其代表的权值来映射成罗马数字
需要注意的是4、9需要单独进行讨论即可
*/
func intToRoman(num int) string {
	var trans = map[int]byte{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}
	var getRoman = func(num *int, carry int) (roman string) {
		var count int
		for *num/carry != 0 {
			count++
			*num -= carry
		}
		if count == 4 || count == 9 {
			if count == 4 {
				roman = string(trans[carry]) + string(trans[carry*5])
			} else {
				roman = string(trans[carry]) + string(trans[carry*10])
			}
		} else {
			if count >= 5 {
				roman = string(trans[carry*5])
				count -= 5
			}
			roman += strings.Repeat(string(trans[carry]), count)
		}

		return roman
	}
	var ans string
	var vec = []int{1000, 100, 10, 1}
	for _, v := range vec {
		ans += getRoman(&num, v)

	}
	return ans
}

func main() {
	intToRoman(99)
}
