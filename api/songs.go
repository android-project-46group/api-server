package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
)

func (server *Server) getAllSongs(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.getAllBlogs")
	defer span.Finish()

	// debug log
	server.logger.Debugf(ctx, "getAllSongs: userAgent", r.UserAgent())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// get group name from query parameters
	group := r.FormValue("gn")

	_, err := server.querier.FindGroupByName(context.Background(), group)
	if err != nil {
		if errors.Unwrap(err) == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			server.logger.Warnf(ctx, "Invalid group name (%s) was passed.", group)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		server.logger.Errorf(ctx, "failed to FindGroupByName: %w", err)
		return
	}
	dr, err := server.querier.GetAllSongs(ctx, group)
	if err != nil {
		server.logger.Errorf(ctx, "failed to get all songs: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// GetSongsResponse
	var res []GetSongsResponse

	for _, r := range dr {

		center := server.querier.GetCenter(ctx, r.Title)
		res = append(res, GetSongsResponse{
			Single: r.Single,
			Title:  r.Title,
			Center: center,
		})
	}

	jRes, _ := json.Marshal(
		map[string]interface{}{
			"songs": res,
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jRes))
}
