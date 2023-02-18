package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/android-project-46group/api-server/db"
	tracer "github.com/opentracing/opentracing-go"
)

func (server *Server) getAllFormations(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.getAllBlogs")
	defer span.Finish()

	// debug log
	server.logger.Debugf(ctx, "getAllFormations: userAgent", r.UserAgent())

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

	dbRes, err := server.querier.GetAllFormations(ctx, group)
	if err != nil {
		server.logger.Errorf(ctx, "failed to get all formations: %w", err)
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := createFormationResponse(dbRes)

	jRes, _ := json.Marshal(
		map[string]interface{}{
			"formations": res,
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jRes))
}

func createFormationResponse(dbRes []db.PositionSongsBind) []FormationResponse {
	if len(dbRes) == 0 {
		return nil
	}

	var res []FormationResponse

	currentSongId := dbRes[0].Song.SongID
	var tmp []Position
	for i, r := range dbRes {

		tmp = append(tmp, Position{
			MemberName: r.Member.NameJa,
			Position:   r.Position.Position,
			IsCenter:   r.Position.IsCenter.Bool,
		})

		if i == len(dbRes)-1 {
			res = append(res, FormationResponse{
				Single:    r.Song.Single,
				Title:     r.Song.Title,
				Positions: tmp,
			})
		} else {
			// 曲ごとにまとめるための処理
			if dbRes[i+1].Song.SongID != currentSongId {
				currentSongId = dbRes[i+1].Song.SongID
				res = append(res, FormationResponse{
					Single:    r.Song.Single,
					Title:     r.Song.Title,
					Positions: tmp,
				})
				tmp = nil
			}
		}
	}
	return res
}
