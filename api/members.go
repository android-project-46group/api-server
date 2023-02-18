package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	models "github.com/android-project-46group/api-server/db/my_models"
	"github.com/android-project-46group/api-server/util"
	"golang.org/x/text/language"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

/*
* /members?gn=sakurazaka
 */
func (server *Server) getAllMembers(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.getAllBlogs")
	defer span.Finish()

	// debug log
	server.logger.Debugf(ctx, "getAllMembers: userAgent", r.UserAgent())

	lang, _ := r.Cookie("lang")
	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(server.matcher, lang.String(), accept)
	locale := tag.String()[0:2]
	key := r.FormValue(queryApiKey)
	// get group name from query parameters
	group := r.FormValue(queryGroupName)
	querySortKey := r.FormValue(querySortKey)
	queryDesc := r.FormValue(queryDesc)

	server.logger.Infof(ctx, "server.getAllMembers, locale %s, key %s, group %s, querySortKey %s, desc", locale, key, group, querySortKey, queryDesc)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

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

	if group != "" {
		_, err := server.querier.FindGroupByName(ctx, group)
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
	}

	// 言語情報を取得する。
	var l *models.Locale
	l, err := server.querier.FindLocaleByName(ctx, locale)
	if err != nil {
		// DB に見つからなかった場合、デフォルトの情報を取得。
		l, _ = server.querier.FindLocaleByName(ctx, "ja")
	}

	sortKey := queryToColumnName(querySortKey)

	// default value is false
	desc := false
	if queryDesc != "" {
		desc, err = strconv.ParseBool(queryDesc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			server.logger.Errorf(ctx, "failed to ParseBool queryDesc: %w", err)
			return
		}
	}

	infos, err := server.querier.GetAllMemberInfos(ctx, group, l.LocaleID, sortKey, desc)
	if err != nil {
		server.logger.Errorf(ctx, "failed to get all memberInfos: %w", err)
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res []MemberInfoResponse
	for _, info := range infos {
		m := MemberInfoResponse{
			MemberId:   info.MemberInfo.MemberID,
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
			"counts":  len(res),
			"members": res,
		},
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}

// queryToColumnName convert query string to SortKey.
func queryToColumnName(query string) util.SortKey {
	switch query {
	case SortKeyID:
		return util.MemberID
	case SortKeyBirthday:
		return util.Birthday
	case SortKeyHeight:
		return util.Height
	default:
		return util.MemberID
	}
}
