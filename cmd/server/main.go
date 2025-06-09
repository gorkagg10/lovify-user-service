package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {
	port := 8082

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	userServer := setupUserServer()

	<-ctx.Done()
}

func setupUserServer(userServer *server)
