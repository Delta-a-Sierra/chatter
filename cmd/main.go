package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Delta-a-Sierra/chatter/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if os.IsExist(err) {
			panic(fmt.Errorf("godotenv.Load: %w", err))
		}
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		if err := app.Start(); err != nil {
			log.Fatal(fmt.Errorf("app.Start: %w", err))
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")
}
