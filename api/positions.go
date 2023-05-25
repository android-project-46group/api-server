package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
)

/*
* /members?gn=sakurazaka
 */
func (server *Server) getPositions(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span, ctx := tracer.StartSpanFromContext(ctx, "api.getPositions")
	defer span.Finish()

	// debug log
	server.logger.Debugf(ctx, "getPositions: userAgent", r.UserAgent())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// get group name from query parameters
	title := r.FormValue("title")

	pMs, err := server.querier.GetPositionFromTitle(ctx, title)
	if err != nil {
		server.logger.Errorf(ctx, "failed to get all positions: %w", err)
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
