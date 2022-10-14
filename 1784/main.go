package main

func checkOnesSegment(s string) bool {
	var zero = false
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero = true
			continue
		}
		if zero {
			if s[i] == '1' {
				return false
			}
		}
	}
	return true
}
func main() {

}
