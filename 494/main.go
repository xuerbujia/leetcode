package main

func findTargetSumWays(nums []int, target int) int {
	var dfs func(idx, target int)
	var ans int
	dfs = func(idx, target int) {
		if idx == len(nums) {
			if target == 0 {
				ans++
			}
			return
		}
		dfs(idx+1, target+nums[idx])
		dfs(idx+1, target-nums[idx])
	}
	return ans
}
func main() {

}
