package main

func removeDuplicateLetters(s string) string {
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	var stack []byte
	var hash = map[byte]bool{}
	for i := 0; i < len(s); i++ {
		for len(stack) > 0 && s[i] < stack[len(stack)-1] && count[stack[len(stack)-1]-'a'] > 0 {
			hash[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		if !hash[s[i]] {
			stack = append(stack, s[i])
			hash[s[i]] = true
		}
		count[s[i]-'a']--
	}

	return string(stack)
}
func main() {

}
