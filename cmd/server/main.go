package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gorkagg10/lovify-user-service/internal/infra/server"
)

func main() {
	_ = 8082

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	_ = setupUserServer()

	<-ctx.Done()
}

func setupUserServer() *server.UserServer {
	return &server.UserServer{}
}
