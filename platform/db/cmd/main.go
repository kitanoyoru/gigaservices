package main

import (
	"context"
	"os/signal"

	app "github.com/kitanoyoru/gigaservices/platform/db/internal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- app.Run()
	}()

	select {
	case err := <-errCh:
		log.Fatalf("Error when runnning app: %+v", err)
	case <-ctx.Done():
		log.Info("Shutting down...")

	}
}
