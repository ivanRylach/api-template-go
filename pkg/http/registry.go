package http

import (
	"api.ivanrylach.github.io/v1/pkg/records"
	"github.com/gin-gonic/gin"
)

func registerHandlers(router *gin.Engine, records *records.Repository) {
	v1 := router.Group("/v1")
	v1.GET("/ping", PingHandler())
	v1.GET("/panic", PanicHandler())

	r := Records{Repo: records}
	v1.POST("/records", r.Create())
	v1.GET("/record/:id", r.Fetch())
}
