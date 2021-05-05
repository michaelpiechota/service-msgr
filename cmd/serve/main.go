package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	health "github.com/InVisionApp/go-health/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/michaelpiechota/service-msgr/serve"
	"go.uber.org/zap"
)

// flag to generate API Router documentation in json format
var printRoutes = flag.Bool("routes", false, "Generate Router Documenation")

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

	// set REST routes for messenger
	r.Route("/messenger", func(r chi.Router) {
		// routes for all users
		// request messages from all senders in the last 30 days
		r.Get("/messages/all", GetAllMessages) // GET /messages/all
		// request messages from all senders limited to last 100 messages
		r.Get("messages/recent", GetRecentMessages) // GET /messages/recent

		// routes that require a user ID to send or retrieve
		r.Route("{userID}", func(r chi.Router) {
			r.Use(UserCtx) // load with request context
			// send message from one sender to one recipient
			r.Post("/send", SendMessage) // POST /{userID}/messenger/send
			// request all messages from specific user in last 30 days
			r.Get("/messages/all", GetAllUserMessages) // GET /{userID}/messages/all
			//request messages from specific user; limited to last 100 messages
			r.Get("/messages/recent", GetRecentMessages) // GET /{userID}/messages/recent
		})
	})

	// generate JSON doc with route info
	if *printRoutes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/service-msgr",
			Intro:       "Generated REST API Docs",
		}))
		return
	}

	server := serve.NewServer(
		config.GetStringOrDefault("PORT", "3000"),
		r,
		logger,
	)

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
