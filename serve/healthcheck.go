package serve

import (
	"net/http"

	health "github.com/InVisionApp/go-health/v2"
	"github.com/InVisionApp/go-health/v2/handlers"
)

// DefaultHealthEndpoint path for grabbing the servers health
const DefaultHealthEndpoint = "healthz"

// DefaultReadyEndpoint path for checking if the server is ready
const DefaultReadyEndpoint = "readyz"

// HealthChecker struct containing working pieces of the server's healthcheck
type HealthChecker struct {
	*health.Health
	handler http.Handler
	logger  *Logger
}

// NewHealthChecker instantiates as new health check
func NewHealthChecker(logger *Logger) *HealthChecker {
	healthcheck := health.New()

	return &HealthChecker{
		Health:  healthcheck,
		handler: handlers.NewJSONHandlerFunc(healthcheck, nil),
		logger:  logger,
	}
}

// Start starts the health check.
func (hc *HealthChecker) Start() error {
	hc.logger.Info("Starting health checker...")

	if err := hc.Health.Start(); err != nil {
		return err
	}

	hc.logger.Info("Health checker started")

	return nil
}

// ServeHTTP implements the http handler interface.
func (hc *HealthChecker) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	hc.handler.ServeHTTP(responseWriter, request)
}
