package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
)

// Handler function to get all groups.
// Must have api-key as a query parameter.
func (server *Server) getAllGroups(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.getAllGroups")
	defer span.Finish()

	// debug log
	server.logger.Debugf(ctx, "getAllGroups: userAgent", r.UserAgent())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if err := server.isApiKeyValid(ctx, key); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			server.logger.Warnf(ctx, "Invalid api key (%s) was passed.", key)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		server.logger.Errorf(ctx, "failed to isApiKeyValid: %w", err)
		return
	}

	gs, err := server.querier.GetAllGroups(ctx)
	if err != nil {
		server.logger.Errorf(ctx, "failed to get all songs: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var data []GroupResponse
	for _, i := range gs {
		data = append(data, GroupResponse{
			Id:   i.GroupID,
			Name: i.GroupName,
		})
	}

	jRes, _ := json.Marshal(
		map[string]interface{}{
			"groups": data,
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jRes))
}
