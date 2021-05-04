package serve

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Server wraps an http server.
type Server struct {
	*http.Server
	logger *Logger
	wg     *sync.WaitGroup
}

const shutdownDuration = time.Second * 30

// NewServer returns a new server.
func NewServer(port string, handler http.Handler, logger *Logger) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: handler,
		},
		logger: logger,
		wg:     &sync.WaitGroup{},
	}
}

// ListenAndServe listens on the server address and serves.
func (s *Server) ListenAndServe() error {
	s.logger.Info("Starting server...")

	go s.awaitShutdownSignal()

	s.RegisterOnShutdown(s.flushLogs)

	s.logger.Info(fmt.Sprintf("Server listening on %s", s.Server.Addr))

	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		s.logger.Error(fmt.Sprintf("Server startup: %v", err))
		return err
	}

	s.wg.Wait()
	s.logger.Info("Server shut down")

	return nil
}

// RegisterOnShutdown registers a handler to run on shutdown.
func (s *Server) RegisterOnShutdown(fn func()) {
	s.wg.Add(1)
	s.Server.RegisterOnShutdown(func() {
		go func() {
			defer s.wg.Done()
			fn()
		}()
	})
}

func (s *Server) awaitShutdownSignal() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	reason := <-sigint

	s.logger.Info("Shutting down server...", zap.String("reason", reason.String()))

	ctx, cancel := context.WithTimeout(context.Background(), shutdownDuration)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		s.logger.Error(fmt.Sprintf("Server shutdown: %v", err))
		return
	}
}

func (s *Server) flushLogs() {
	s.logger.Info("Flushing logs...")
	s.logger.Flush()
	s.logger.Info("Logs flushed")
}
