package main

import (
	"net/http/cgi"
	"os"

	"github.com/gorilla/mux"

	"web/api"
	"web/db"
)


func main() {
	// DB の初期設定をする
	_db, err := db.DbInit()
	if err != nil {
		// panic("")
	}
	defer _db.Close()
		
	r := mux.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")
	r.Path(rootPath + "/members").
		Queries("gn", "{gn}").
		HandlerFunc(api.GetAllMembers).
		Methods("GET")

	cgi.Serve(r)
}