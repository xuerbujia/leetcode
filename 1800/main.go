package main

func maxAscendingSum(nums []int) int {
	var ans int
	var pre int
	var m int
	for _, v := range nums {
		if v > pre {
			m += v
			ans = max(ans, m)
		} else {
			m = v
		}
		pre = v
	}
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

}
