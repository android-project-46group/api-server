package api

import (
	"fmt"
	"net/http"
	"encoding/json"

	"web/db"
)


func GetAllSongs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	key := r.FormValue("key")
	
	if !IsApiKeyValid(key) {
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

	dr, err := db.GetAllSongs(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// GetSongsResponse
	var res []GetSongsResponse

	for _, r := range dr {

		center := db.GetCenter(r.Title)
		res = append(res, GetSongsResponse{
			Single: r.Single,
			Title: r.Title,
			Center: center,
		})
	}

	jRes, _ := json.Marshal(
		map[string]interface{} {
			"songs": res,
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jRes))
}