package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
)

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		span, ctx := tracer.StartSpanFromContext(ctx, "api.authMiddleware")
		defer span.Finish()

		key := r.FormValue(queryApiKey)
		if err := s.isApiKeyValid(ctx, key); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				s.logger.Warnf(ctx, "Invalid api key (%s) was passed.", key)
				http.Error(w, "invalid api key", http.StatusUnauthorized)

				return
			}

			s.logger.Errorf(ctx, "failed to isApiKeyValid: %w", err)
			http.Error(w, fmt.Sprintf("server error: %v", err), http.StatusInternalServerError)

			return
		}

		next.ServeHTTP(w, r)
	})
}
