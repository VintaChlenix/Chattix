package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handlers "Chattix/internal/chat/delivery/http"
	"golang.org/x/sync/errgroup"
)

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	chattixDelivery := handlers.NewChat()
	slog.Info("app initialized")

	http.Handle("/", chattixDelivery.Handler())

	srv := http.Server{
		Addr: "0.0.0.0:80",
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		slog.Info("Starting server")
		lerr := srv.ListenAndServe()
		if errors.Is(lerr, http.ErrServerClosed) {
			return nil
		}

		return lerr
	})

	eg.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(ctx)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	slog.Info("Shutdown app")
	return nil
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
