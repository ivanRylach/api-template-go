package api_v1

import "github.com/gin-gonic/gin"

func PanicHandler() func(c *gin.Context) {
    return func(c *gin.Context) {
        panic("Oh no!")
    }
}