package api

import (
	"database/sql"
	"encoding/json"
	"errors"
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

	if err := server.isApiKeyValid(key); err != nil {
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

	// 言語情報を取得する。
	var l *models.Locale
	l, err = server.querier.FindLocaleByName(locale)
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
