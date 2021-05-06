package messenger

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// MessageCtx is used to load an Message object from the URL parameters
func MessageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var message *Message
		var err error

		if userID := chi.URLParam(r, "userID"); userID != "" {
			message, err = dbGetMessage(userID)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "message", message)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
