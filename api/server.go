package api

import (
	"net/http"
	"net/http/cgi"
	"net/url"
	"os"

	"golang.org/x/text/language"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"

	"github.com/android-project-46group/api-server/db"
	"github.com/android-project-46group/api-server/repository/grpc"
	"github.com/android-project-46group/api-server/util"
)

// Server serves HTTP requests for this service.
type Server struct {
	config     util.Config
	querier    db.Querier
	router     *muxtrace.Router
	matcher    language.Matcher
	logger     util.Logger
	grpcClient grpc.GrpcClient
}

// Create a new server from the given config file.
func NewServer(
	config util.Config, querier db.Querier, matcher language.Matcher,
	logger util.Logger, grpcClient grpc.GrpcClient,
) (*Server, error) {

	server := &Server{
		config:     config,
		querier:    querier,
		matcher:    matcher,
		logger:     logger,
		grpcClient: grpcClient,
	}

	err := server.setupRouter()
	return server, err
}

func (server *Server) setupRouter() error {
	r := muxtrace.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")
	var err error
	if server.config.URLPrefix != "" {
		rootPath, err = url.JoinPath("/", rootPath, server.config.URLPrefix)
		if err != nil {
			return err
		}
	}

	r.Path(rootPath + "/health").
		HandlerFunc(server.Health).
		Methods("GET")

	r.Path(rootPath + "/health/grpc").
		HandlerFunc(server.healthGrpc).
		Methods("GET")

	authed := r.PathPrefix("").Subrouter()
	authed.Use(server.authMiddleware)

	authed.Path(rootPath+"/groups").
		Queries("key", "{key}").
		HandlerFunc(server.getAllGroups).
		Methods("GET")

	authed.Path(rootPath+"/members").
		Queries("key", "{key}").
		HandlerFunc(server.getAllMembers).
		Methods("GET")

	authed.Path(rootPath+"/blogs").
		Queries("key", "{key}").
		HandlerFunc(server.getAllBlogs).
		Methods("GET")

	authed.Path(rootPath+"/songs").
		Queries("key", "{key}").
		HandlerFunc(server.getAllSongs).
		Methods("GET")

	authed.Path(rootPath+"/positions").
		Queries("title", "{title}").
		Queries("key", "{key}").
		HandlerFunc(server.getPositions).
		Methods("GET")

	authed.Path(rootPath+"/formations").
		Queries("key", "{key}").
		HandlerFunc(server.getAllFormations).
		Methods("GET")

	authed.Path(rootPath + "/members/download").
		HandlerFunc(server.downloadMembers).
		Methods("GET")

	server.router = r

	return nil
}

// start server depending on the cgi serve or not
func (server *Server) Start() error {

	if server.config.IsCGI {
		return cgi.Serve(server.router)
	}

	return http.ListenAndServe(server.config.ServerAddress, server.router)
}

// Serve as a CGI program.
func (server *Server) CGI() {
	cgi.Serve(server.router)
}
