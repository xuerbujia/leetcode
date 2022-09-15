package main

import "fmt"

type coder interface {
	code()
	debug()
	getLanguage() string
}

type Gopher struct {
	language string
}

func (c Gopher) code() {
	c.language = "PHP"
	fmt.Println("I am coding", c.language)
}
func (c Gopher) debug() {
	fmt.Println("I am debug", c.language)
}
func (c Gopher) getLanguage() string {
	return c.language
}
func main() {
	var _ coder = Gopher{}
}
