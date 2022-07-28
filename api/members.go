package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
* /members?gn=sakurazaka
 */
func (server *Server) getAllMembers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if !server.isApiKeyValid(key) {
		fmt.Printf("getAllMembers: access with invalid api key")
		// return error message
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, ErrorJson("No valid api key"))
		return
	}

	// get group name from query parameters
	group := r.FormValue("gn")

	if !server.querier.ExistGroup(group) {
		fmt.Printf("getAllFormations: access to invalid group name")
		// return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("Error while"))
		return
	}
	infos, err := server.querier.GetAllMemberInfos(group)
	if err != nil {
		fmt.Printf("getAllFormations: %v", err)
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res []MemberInfoResponse
	for _, info := range infos {
		m := MemberInfoResponse{
			MemberId:   info.Member.MemberID,
			MemberName: info.Member.NameJa,
			Birthday:   info.MemberInfo.Birthday,
			Height:     info.MemberInfo.Height,
			BloodType:  info.MemberInfo.BloodType,
			Generation: info.MemberInfo.Generation,
			BlogURL:    info.MemberInfo.BlogURL.String,
			ImgURL:     info.MemberInfo.ImgURL.String,
		}
		res = append(res, m)
	}

	// make json for http response
	jsonRes, _ := json.Marshal(
		map[string]interface{}{
			"members": res,
		},
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}
