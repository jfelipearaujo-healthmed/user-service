package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/config"
	"github.com/jfelipearaujo-healthmed/user-service/internal/core/infrastructure/shared/logger"
	"github.com/jfelipearaujo-healthmed/user-service/internal/external/http/server"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	time.Local, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	if err := godotenv.Load(); err != nil {
		slog.ErrorContext(ctx, "error loading .env file on root folder", "error", err)

		if err := godotenv.Load("../../.env"); err != nil {
			slog.ErrorContext(ctx, "error loading .env file", "error", err)
			panic(err)
		}
	}

	config, err := config.LoadFromEnv(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error loading config from env", "error", err)
		panic(err)
	}

	logger.SetupLog(config)

	server, err := server.NewServer(ctx, config)
	if err != nil {
		slog.ErrorContext(ctx, "error creating server", "error", err)
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go startHttpServer(ctx, &wg, server)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	cancel()
	wg.Wait()

	slog.InfoContext(ctx, "graceful shutdown completed ✅")
}

func startHttpServer(ctx context.Context, wg *sync.WaitGroup, server *server.Server) {
	defer wg.Done()

	httpServer := server.GetServer()

	go func() {
		slog.InfoContext(ctx, "🚀 Server started", "address", httpServer.Addr)
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "http server error", "error", err)
			panic(err)
		}
		slog.InfoContext(ctx, "http server stopped serving requests")
	}()

	<-ctx.Done()

	shutdownCtx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdown()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		slog.ErrorContext(ctx, "error while trying to shutdown the server", "error", err)
	}

	if err := server.DbService.Close(ctx); err != nil {
		slog.ErrorContext(ctx, "error while trying to close the database connection", "error", err)
	}

	if err := server.Cache.Close(ctx); err != nil {
		slog.ErrorContext(ctx, "error while trying to close the cache connection", "error", err)
	}
}
