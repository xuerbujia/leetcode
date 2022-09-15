package main

import "strings"

func reorderSpaces(text string) string {
	var words []string
	var word []byte
	var wordLen int
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' || i == len(text)-1 {
			if len(word) != 0 {
				wordLen += len(word)
				words = append(words, string(word))
				word = make([]byte, 0)
			}
			continue
		}
		word = append(word, text[i])
	}
	space := len(text) - wordLen
	repeat := strings.Repeat(" ", space/(len(words)-1))
	join := strings.Join(words, repeat)
	repeat = strings.Repeat(" ", len(text)%(len(words)-1))
	return join + repeat
}
func main() {

}
