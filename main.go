package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/multimoml/stats/internal/server"
)

// @title Statistics API
// @version 1.0.0
// @host localhost:6004
// @BasePath /stats
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go server.Run(ctx)

	<-sigChan
	cancel()
}
