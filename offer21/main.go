package main

func exchange(nums []int) []int {
	var ans []int
	var temp []int
	for _, v := range nums {
		if v%2 == 0 {
			temp = append(temp, v)
		} else {
			ans = append(ans, v)
		}
	}
	ans = append(ans, temp...)
	return ans
}
func main() {

}
