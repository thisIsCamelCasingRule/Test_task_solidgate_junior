package main

import (
	"Test_task_solidgate_junior/cmd/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.NewServer()

	err := srv.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
