package api

import (
	"net/http"
	"net/http/cgi"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/text/language"

	"github.com/android-project-46group/api-server/db"
	"github.com/android-project-46group/api-server/util"
)

// Server serves HTTP requests for this service.
type Server struct {
	config  util.Config
	querier db.Querier
	router  *mux.Router
	matcher language.Matcher
}

// Create a new server from the given config file.
func NewServer(config util.Config, querier db.Querier, matcher language.Matcher) (*Server, error) {

	server := &Server{
		config:  config,
		querier: querier,
		matcher: matcher,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	r := mux.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")
	r.Path(rootPath+"/members").
		Queries("gn", "{gn}").
		HandlerFunc(server.getAllMembers).
		Methods("GET")

	r.Path(rootPath+"/blogs").
		Queries("gn", "{gn}").
		HandlerFunc(server.getAllBlogs).
		Methods("GET")

	r.Path(rootPath+"/songs").
		Queries("gn", "{gn}").
		HandlerFunc(server.getAllSongs).
		Methods("GET")

	r.Path(rootPath+"/positions").
		Queries("title", "{title}").
		HandlerFunc(server.getPositions).
		Methods("GET")

	r.Path(rootPath+"/formations").
		Queries("gn", "{gn}").
		HandlerFunc(server.getAllFormations).
		Methods("GET")

	server.router = r
}

// start server depending on the cgi serve or not
func (server *Server) Start() error {

	if server.config.IsCGI {
		return cgi.Serve(server.router)
	}

	return http.ListenAndServe(server.config.ServerAddress, server.router)
}
