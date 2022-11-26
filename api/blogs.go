package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

/*
* /blogs?gn=sakurazaka&key=xxxyyyzzzhogehoge
 */
func (server *Server) getAllBlogs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if !server.isApiKeyValid(key) {
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

	blogs, err := server.querier.GetAllBlogs(group)
	if err != nil {
		// db error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res []BlogResponse
	for _, b := range blogs {
		m := BlogResponse{
			MemberName:    b.Member.NameJa,
			BlogURL:       b.Blog.BlogURL,
			LastBlogImg:   b.Blog.LastBlogImg,
			LastUpdatedAt: b.Blog.LastUpdatedAt,
		}
		res = append(res, m)
	}

	// make json for http response
	jsonRes, _ := json.Marshal(
		map[string]interface{}{
			"blogs": res,
		},
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}
