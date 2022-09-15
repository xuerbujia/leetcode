package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var ans int
	for i := 0; i < len(startTime); i++ {
		if queryTime > startTime[i] && queryTime < endTime[i] {
			ans++
		}
	}
	return ans
}
func main() {

}
