package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (server *Server) Health(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.Health")
	defer span.Finish()

	jsonRes, _ := json.Marshal(
		map[string]string{
			"status": "healthy",
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
	server.logger.Debug(ctx, "Health checked")
}
