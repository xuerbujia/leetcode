package main

import (
	"fmt"
	"math"
	"strings"
)

func strToInt(str string) int {
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	symbol := 1
	if s[0] < '0' || s[0] > '9' {
		if s[0] == '+' || s[0] == '-' {
			if s[0] == '-' {
				symbol = -1
			}
			s = s[1:]
		} else {
			return 0
		}
	}
	var ans int

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		ans *= 10
		ans += int(s[i] - '0')
		if ans*symbol > math.MaxInt32 {
			return math.MaxInt32
		}
		if ans*symbol < math.MinInt32 {
			return math.MinInt32
		}
	}

	return ans * symbol
}
func main() {
	fmt.Println(strToInt("9223372036854775808"))
}
