package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
* /members?gn=sakurazaka
 */
func (server *Server) getPositions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if !server.isApiKeyValid(key) {
		// return error message
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, ErrorJson("No valid api key"))
		return
	}

	// get group name from query parameters
	title := r.FormValue("title")

	pMs, err := server.querier.GetPositionFromTitle(title)
	if len(pMs) == 0 {
		// db error
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("Invalid title"))
		return
	}
	if err != nil {
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res []GetPositionsResponse
	for _, pm := range pMs {
		m := GetPositionsResponse{
			MemberName: pm.Member.NameJa,
			ImgURL:     pm.MemberInfo.ImgURL.String,
			Position:   pm.Position.Position,
			IsCenter:   pm.Position.IsCenter.Bool,
		}
		res = append(res, m)
	}

	// make json for http response
	jsonRes, _ := json.Marshal(
		map[string]interface{}{
			"positions": res,
		},
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}
