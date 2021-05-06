package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	health "github.com/InVisionApp/go-health/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
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
		r.Post("/", messenger.CreateMessage) // POST /messages/{userID}
		// get all messages from all senders limited to last 100 messages
		r.Get("/search", messenger.SearchMessages) // GET /messages/search

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(messenger.MessageCtx)            // use request context
			r.Get("/", messenger.GetMessage)       // GET /messages/{userID}
			r.Delete("/", messenger.DeleteMessage) // DELETE /messages/{userID}
		})

	})

	// flag to generate API Router documentation in json format
	var printRoutes = flag.Bool("printRoutes", false, "Generate Router Documenation")
	// generate JSON doc with route info
	if *printRoutes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/service-msgr",
			Intro:       "Generated REST API Docs",
		}))
		return
	}

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
