package main

import (
    "api.ivanrylach.github.io/v1/pkg/http"
    "api.ivanrylach.github.io/v1/pkg/util"
    "context"
    "go.uber.org/zap"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    util.ConfigureLogging()
    httpServer := http.NewServer()
    go httpServer.Start()

    // Wait for interrupt signal to gracefully shutdown the server with
    // a timeout of 5 seconds.
    quit := make(chan os.Signal)
    // kill (no param) default send syscall.SIGTERM
    // kill -2 is syscall.SIGINT
    // kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    zap.S().Info("Shutting down server...")

    // The context is used to inform the server it has 5 seconds to finish
    // the request it is currently handling
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    httpServer.Stop(&ctx)

    zap.S().Info("Server stopped...")
}
