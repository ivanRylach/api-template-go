package http

import "github.com/gin-gonic/gin"

func PanicHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		panic("Oh no!")
	}
}
