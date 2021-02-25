package main

import (
	"api.ivanrylach.github.io/v1/pkg/http"
	"api.ivanrylach.github.io/v1/pkg/mongodb"
	"api.ivanrylach.github.io/v1/pkg/records"
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
	mongo := mongodb.NewClient("mongodb://root:password123@localhost:27017")
	recordsRepo := records.NewRepository(mongo)

	httpServer := http.NewServer(recordsRepo)
	go httpServer.Start()

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongo.Shutdown(ctx)
	httpServer.Stop(ctx)

}
