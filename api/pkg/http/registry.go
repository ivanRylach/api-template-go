package http

import (
	"github.com/gin-gonic/gin"
	"ivanrylach.github.io/api/v1/pkg/records"
)

func registerHandlers(router *gin.Engine, records *records.Repository) {
	v1 := router.Group("/v1")
	v1.GET("/ping", PingHandler())
	v1.GET("/panic", PanicHandler())

	r := Records{Repo: records}
	v1.POST("/records", r.Create())
	v1.GET("/record/:id", r.Fetch())
}
