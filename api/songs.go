package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (server *Server) getAllSongs(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.getAllBlogs")
	defer span.Finish()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if err := server.isApiKeyValid(ctx, key); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, ErrorJson("No valid api key"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, ErrorJson("Error while reading api key from DB"))
		return
	}

	// get group name from query parameters
	group := r.FormValue("gn")

	_, err := server.querier.FindGroupByName(context.Background(), group)
	if err != nil {
		if errors.Unwrap(err) == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, ErrorJson("Invalid group name was passed."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, ErrorJson("Error while reading group from DB"))
		return
	}
	dr, err := server.querier.GetAllSongs(ctx, group)
	if err != nil {
		fmt.Printf("getAllSongs: %v", err)
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
