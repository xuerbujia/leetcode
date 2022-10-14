package main

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var count []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count = append(count, i)
		}
	}
	return len(count) == 0 || (len(count) == 2 && s1[count[0]] == s2[count[1]] && s1[count[1]] == s2[count[0]])
}
func main() {

}
