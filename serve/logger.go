package serve

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// Logger struct to contain logger data and functionality
type Logger struct {
	*zap.Logger
}

type AppLogger interface {
	Error(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Flush()
}

const (
	// LoggerTypeTest test log level
	LoggerTypeTest = "TEST"
	// LoggerTypeDevelopment development log level
	LoggerTypeDevelopment = "DEVELOPMENT"
)

// NewLogger creates a new Chi logger with Zap. Will create the log level based on the parameter sent
func NewLogger(loggerType string) *Logger {
	var logger *zap.Logger
	var err error

	switch loggerType {
	case LoggerTypeTest:
		logger = zap.NewNop()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	return &Logger{logger}
}

// Print prints the args sent to the logger
func (l *Logger) Print(args ...interface{}) {
	fmt.Print(args...)
}

// Flush flushes all the logs to the io of choice.
// Should be used when shutting down the server to ensure all logs are printed
func (l *Logger) Flush() {
	if err := l.Sync(); err != nil {
		return
	}
}

// NewRequestLoggerMiddleware middlware for Chi router to log the status of a request
func NewRequestLoggerMiddleware(l *Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			start := time.Now()

			defer func() {
				l.Info("Served",
					zap.String("protocol", r.Proto),
					zap.String("path", r.URL.Path),
					zap.Duration("responseTime", time.Since(start)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
					zap.String("reqId", middleware.GetReqID(r.Context())),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
