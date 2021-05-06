package main

import (
	"fmt"
	"os"
	"path"

	health "github.com/InVisionApp/go-health/v2"
	"github.com/go-chi/chi/v5"
	"github.com/michaelpiechota/service-msgr/messenger"
	"github.com/michaelpiechota/service-msgr/serve"
	"go.uber.org/zap"
)

func main() {
	// grab env vars
	config := serve.NewConfig(os.Getenv("ENV_FILE"))
	// setup structured logging using zap
	logger := serve.NewLogger(config.GetStringOrDefault("LOGGER", "DEVELOPMENT"))
	// setup healthcheck for http server
	healthcheck := serve.NewHealthChecker(logger)
	// add more healthchecks here
	healthchecks := []*health.Config{}

	if err := healthcheck.AddChecks(healthchecks); err != nil {
		logger.Fatal("Failed to add health checks: %v", zap.Error(err))
	}

	// create new router from serve package
	r := serve.NewRouter(logger)
	r.Handle(healthCheckEndpoint(config), healthcheck)
	r.Handle(readyCheckEndpoint(config), healthcheck)

	// set routes
	r.Route("/messages", func(r chi.Router) {
		r.With(messenger.Paginate).Get("/", messenger.ListMessages)
		// recent messages can be requested from all senders - with a limit of 100
		// messages or all messages in last 30 days.
		r.Get("/search", messenger.SearchMessages) // GET /messages/search

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(messenger.MessageCtx) // use request context
			// recent messages can be requested for a recipient from a specific sender
			// with a limit of 100 messages or all messages in last 30 days.
			r.Get("/", messenger.GetMessage)       // GET /messages/{userID}
			r.Delete("/", messenger.DeleteMessage) // DELETE /messages/{userID}
		})

		r.Route("/create/{userID}", func(r chi.Router) {
			r.Use(messenger.NewMessageCtx)       // use request context
			r.Post("/", messenger.CreateMessage) // POST /messages/{userID} where {userID} is the recipient's ID
		})

	})

	// start http server
	server := serve.NewServer(
		config.GetStringOrDefault("PORT", "3000"),
		r,
		logger,
	)

	// start healthcheck
	if err := healthcheck.Start(); err != nil {
		logger.Fatal("Failed to start health check: %v", zap.Error(err))
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func healthCheckEndpoint(config *serve.Config) string {
	endpoint := config.GetStringOrDefault("HEALTH_CHECK_ENDPOINT", serve.DefaultHealthEndpoint)
	return path.Clean(fmt.Sprintf("/%s", endpoint))
}

func readyCheckEndpoint(config *serve.Config) string {
	endpoint := config.GetStringOrDefault("READY_CHECK_ENDPOINT", serve.DefaultReadyEndpoint)
	return path.Clean(fmt.Sprintf("/%s", endpoint))
}
