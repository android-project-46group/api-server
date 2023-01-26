package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (server *Server) Health(w http.ResponseWriter, r *http.Request) {
	jsonRes, _ := json.Marshal(
		map[string]string{
			"status": "healthy",
		},
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonRes))
}
