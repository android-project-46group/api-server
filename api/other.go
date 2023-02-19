package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
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

func (server *Server) healthGrpc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.healthGrpc")
	defer span.Finish()

	msg, err := server.grpcClient.Health(ctx)
	if err != nil {
		server.logger.Warn(ctx, "failed do health check")

		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(msg))
	w.WriteHeader(http.StatusOK)
}
