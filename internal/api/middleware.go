package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func RequestLogger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId, _ := uuid.NewV7()

			logger.Info(
				r.Method,
				r.URL.Path,
				"request_id",
				reqId.String(),
			)

			ctx := context.WithValue(r.Context(), "requestId", reqId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
