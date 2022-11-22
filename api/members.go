package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/android-project-46group/api-server/db/my_models"
	"golang.org/x/text/language"
)

/*
* /members?gn=sakurazaka
 */
func (server *Server) getAllMembers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	lang, _ := r.Cookie("lang")
	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(server.matcher, lang.String(), accept)
	locale := tag.String()[0:2]

	key := r.FormValue("key")

	if !server.isApiKeyValid(key) {
		// return error message
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, ErrorJson("No valid api key"))
		return
	}

	// get group name from query parameters
	group := r.FormValue("gn")

	if !server.querier.ExistGroup(group) {
		// return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("Error while"))
		return
	}

	// 言語情報を取得する。
	var l *models.Locale
	l, err := server.querier.FindLocaleByName(locale)
	if err != nil {
		// DB に見つからなかった場合、デフォルトの情報を取得。
		l, _ = server.querier.FindLocaleByName("ja")
	}

	infos, err := server.querier.GetAllMemberInfos(group, l.LocaleID)
	if err != nil {
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
