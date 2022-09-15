package main

func constructArray(n int, k int) []int {
	var ans []int
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	var left, right = n - k, n
	for left < right {
		ans = append(ans, left, right)
		left++
		right--
	}
	return ans
}
func main() {

}
