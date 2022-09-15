package main

func maxScore(s string) int {
	var zero, one = make([]int, len(s)), make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i-1] == '0' {
			zero[i] = zero[i-1] + 1
		} else {
			zero[i] = zero[i-1]
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i+1] == '1' {
			one[i] = one[i+1] + 1
		} else {
			one[i] = one[i+1]
		}
	}
	zero = zero[1:]
	one = one[:len(one)-1]
	var max int
	for i := range zero {
		if zero[i]+one[i] > max {
			max = zero[i] + one[i]
		}
	}

	return max
}
func main() {
	maxScore("011101")
}
