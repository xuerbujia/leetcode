package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func minOperations(logs []string) int {
	var stack []string
	for _, v := range logs {

		if v[:len(v)-1] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v[:len(v)-1] == "." {
			continue
		} else {
			stack = append(stack, v[:len(v)-1])
		}
	}
	return len(stack)
}
func Run() {
	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"沙口": "沙口"})
	})
	r.Run(":8080")
}
func main() {
	Run()
}
