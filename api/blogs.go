package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/android-project-46group/api-server/db"
)

/*
* /blogs?gn=sakurazaka&key=xxxyyyzzzhogehoge
 */
func GetAllBlogs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	key := r.FormValue("key")

	if !IsApiKeyValid(key) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, ErrorJson("No valid api key"))
		return
	}

	// get group name from query parameters
	group := r.FormValue("gn")

	if !db.ExistGroup(group) {
		// return error message
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ErrorJson("Error: No valid group name"))
		return
	}
	blogs, err := db.GetAllBlogs(group)
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
