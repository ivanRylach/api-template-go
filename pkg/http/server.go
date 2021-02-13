package http

import (
    "context"
    "github.com/gin-contrib/zap"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "net/http"
    "time"
)

type Server struct {
    Server *http.Server
}

func NewServer() *Server {
    router := gin.New()

    router.Use(ginzap.Ginzap(zap.S().Desugar(), time.RFC822, false))
    router.Use(ginzap.RecoveryWithZap(zap.S().Desugar(), false))

    registerHandlers(router)

    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    return &Server{Server: srv}
}

func (s *Server) Start() {
    zap.S().Info("Starting HTTP server...")
    if err := s.Server.ListenAndServe(); err != nil {
        zap.S().Error(err)
    }
}

func (s *Server) Stop(ctx *context.Context) {
    zap.S().Info("Stopping HTTP server...")
    if err := s.Server.Shutdown(*ctx); err != nil {
        zap.S().Panic(err)
    }
}