package api_v1

import "github.com/gin-gonic/gin"

func PingHandler() func(c *gin.Context) {
    return func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    }
}

