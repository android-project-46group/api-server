package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/android-project-46group/api-server/db"
)

func (server *Server) getAllFormations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if !server.isApiKeyValid(key) {
		fmt.Printf("getAllFormations: access with invalid api key")
		// return error message
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, ErrorJson("No valid api key"))
		return
	}

	// get group name from query parameters
	group := r.FormValue("gn")

	_, err := server.querier.FindGroupByName(group)
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

	dbRes, err := server.querier.GetAllFormations(group)
	if err != nil {
		fmt.Printf("getAllFormations: %v", err)
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
