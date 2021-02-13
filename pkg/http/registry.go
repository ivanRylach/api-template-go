package http

import (
    "api.ivanrylach.github.io/v1/pkg/http/api_v1"
    "github.com/gin-gonic/gin"
)

func registerHandlers(router *gin.Engine) {
    v1 := router.Group("/v1")
    v1.GET("/ping", api_v1.PingHandler())
    v1.GET("/panic", api_v1.PanicHandler())
}
