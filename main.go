package main

import (
	"log"

	"github.com/android-project-46group/api-server/api"
	"github.com/android-project-46group/api-server/db"
	"github.com/android-project-46group/api-server/util"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("saka-api"),
	)
	defer tracer.Stop()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	querier, err := db.NewQuerier(config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer querier.DB.Close()

	matcher := util.NewMatcher()

	// DI to server
	server, err := api.NewServer(config, querier, matcher)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	if config.IsCGI {
		server.CGI()
	} else {
		err = server.Start()
		if err != nil {
			log.Fatal("cannot start server:", err)
		}
	}
}
