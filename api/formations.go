package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/android-project-46group/api-server/db"
)

func (server *Server) getAllFormations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if !server.isApiKeyValid(key) {
		// return error message
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, ErrorJson("No valid api key"))
		return
	}

	// get group name from query parameters
	group := r.FormValue("gn")

	if !db.ExistGroup(group) {
		// return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("Invalid group name"))
		return
	}

	dbRes, err := db.GetAllFormations(group)
	if err != nil {
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res []FormationResponse
	if len(dbRes) == 0 {
		// return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("No formation was registerd for this group"))
		return
	}

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

	jRes, _ := json.Marshal(
		map[string]interface{}{
			"formations": res,
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jRes))
}
