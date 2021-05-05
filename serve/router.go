package serve

import (
	"compress/flate"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// Router chi mux router struct
type Router struct {
	*chi.Mux
}

// NewRouter returns a new router.
func NewRouter(logger *Logger) *Router {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Use the default compressor settings (good mix of speed vs. size)
	compressor := middleware.NewCompressor(flate.DefaultCompression)
	r.Use(compressor.Handler)

	r.Use(NewRequestLoggerMiddleware(logger))

	return &Router{r}
}
