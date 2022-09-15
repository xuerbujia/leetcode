package main

import "fmt"

func stringMatching(words []string) []string {
	var hash = map[string]bool{}
	for _, v := range words {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j <= len(v); j++ {
				if v[i:j] == v {
					continue
				}
				hash[v[i:j]] = true
			}
		}
	}
	//fmt.Println(hash)
	//fmt.Println(hash["mass"])
	var ans []string
	for _, v := range words {
		if hash[v] {
			ans = append(ans, v)
		}
	}
	return ans
}
func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words))
}
