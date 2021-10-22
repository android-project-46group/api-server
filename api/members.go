package api

import (
	"fmt"
	"net/http"
	"encoding/json"

	"web/db"
)

/*
* /members?gn=sakurazaka
*/
func GetAllMembers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// get group name from query parameters
	group := r.FormValue("gn")

	if !db.ExistGroup(group) {
		// return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("Error while"))
		return
	}
	infos, err := db.GetAllMemberInfos(group)
	if err != nil {
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res []MemberInfoResponse
	for _, info := range infos {
		m := MemberInfoResponse{
			MemberId: info.Member.MemberID,
			MemberName: info.Member.NameJa,
			Birthday: info.MemberInfo.Birthday,
			Height: info.MemberInfo.Height,
			BloodType: info.MemberInfo.BloodType,
			Generation: info.MemberInfo.Generation,
			BlogURL: info.MemberInfo.BlogURL.String,
			ImgURL: info.MemberInfo.ImgURL.String,
		}
		res = append(res, m)
	}

	// make json for http response
	jsonRes, _ := json.Marshal(
		map[string]interface{} {
			"members": res,
		},
	)
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}
