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

	r.Path(rootPath + "/blogs").
		Queries("gn", "{gn}").
		HandlerFunc(api.GetAllBlogs).
		Methods("GET")

	r.Path(rootPath + "/formations").
		Queries("gn", "{gn}").
		HandlerFunc(api.GetAllFormations).
		Methods("GET")

	r.Path(rootPath + "/songs").
		Queries("gn", "{gn}").
		HandlerFunc(api.GetAllSongs).
		Methods("GET")

	r.Path(rootPath + "/positions").
		Queries("title", "{title}").
		HandlerFunc(api.GetPositions).
		Methods("GET")
	
	cgi.Serve(r)
}
