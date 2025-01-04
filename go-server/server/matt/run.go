package matt

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/0xBradock/go-srvr/internal/config"
	"github.com/0xBradock/go-srvr/internal/user"
)

func Run(ctx context.Context, w io.Writer, getenv func(string) string) error {
	config := &config.Config{
		Port:            "8888",
		Host:            "0.0.0.0",
		ShutdownTimeout: time.Duration(10),
	}

	logger := slog.New(slog.NewJSONHandler(w, nil))

	userStore := user.MockUser{}

	// TODO: Not sure if logger and config is instantiated here or before
	srv := NewServer(logger, config, userStore)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	go func() {
		slog.Info("starting server", "port", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", "err", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, config.ShutdownTimeout*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("failed to shutdown server", "err", err)
		}
	}()

	wg.Wait()

	return nil
}
