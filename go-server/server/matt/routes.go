package matt

import (
	"log/slog"
	"net/http"

	"github.com/0xBradock/go-srvr/internal/health"
	"github.com/0xBradock/go-srvr/internal/user"
	"github.com/0xBradock/go-srvr/server/matt/middleware"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	// config *Config,
	userStore user.UserStore,
) {
	// Middlewares
	auth := middleware.AuthForRoles

	// Roles
	admin := user.Admin
	mangr := user.Manager

	// Routes
	mux.Handle("/health", health.HandleHealth(logger))
	mux.Handle("POST /user", auth(user.HandleUser(logger, userStore), admin, mangr))
}
