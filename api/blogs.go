package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
)

/*
* /blogs?gn=sakurazaka&key=xxxyyyzzzhogehoge
 */
func (server *Server) getAllBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	span, ctx := tracer.StartSpanFromContext(ctx, "api.getAllBlogs")
	defer span.Finish()

	// debug log
	server.logger.Debugf(ctx, "getAllBlogs: userAgent", r.UserAgent())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// get group name from query parameters
	group := r.FormValue("gn")

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

	blogs, err := server.querier.GetAllBlogs(ctx, group)
	if err != nil {
		server.logger.Errorf(ctx, "failed to get all blogs: %w", err)
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
			"counts": len(res),
			"blogs":  res,
		},
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}
