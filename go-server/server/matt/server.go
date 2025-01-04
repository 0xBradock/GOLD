package matt

import (
	"log/slog"
	"net/http"

	"github.com/0xBradock/go-srvr/internal/config"
	"github.com/0xBradock/go-srvr/internal/user"
)

// NewServer is responsible for all top-level functionality
// common to all endpoints: CORS, auth, logging, etc.
//
// Passing nil when testing to undesired dependencies
func NewServer(
	logger *slog.Logger,
	config *config.Config,
	userStore user.UserStore,
	// productStore *productStore,
) http.Handler {
	mux := http.NewServeMux()

	// attach routes to the server
	addRoutes(
		mux,
		logger,
		// config,
		userStore,
	)

	// attach middlewares to the server
	var handler http.Handler = mux
	// handler = traceMW(handler)
	// handler = loggerMW(handler)
	// handler = authMW(handler)
	// handler = corsMW(handler)

	return handler
}
