package api

import (
	"backentrymiddle/internal/logger"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func Logging(log *slog.Logger) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log = log.With(
				slog.String("ip", r.RemoteAddr),
				slog.String("URL", r.URL.Path),
			)

			ctx := logger.NewContext(r.Context(), log)

			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}
