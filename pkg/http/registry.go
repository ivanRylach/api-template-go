package http

import (
    "api.ivanrylach.github.io/v1/pkg/http/api_v1"
    "api.ivanrylach.github.io/v1/pkg/records"
    "github.com/gin-gonic/gin"
)

func registerHandlers(router *gin.Engine, records *records.Repository) {
    v1 := router.Group("/v1")
    v1.GET("/ping", api_v1.PingHandler())
    v1.GET("/panic", api_v1.PanicHandler())

    r := api_v1.Records{Repo: records}
    v1.POST("/records", r.Create())
    v1.GET("/record/:id", r.Fetch())
}
