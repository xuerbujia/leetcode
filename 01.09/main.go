package main

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i:]+s1[:i] == s2 {
			return true
		}

	}
	return false
}
func main() {
	isFlipedString("", "")
}
