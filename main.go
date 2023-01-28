package main

import (
	"context"
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

	logger, closeFunc, err := util.NewZapLogger(config.LogPath, config.Host, config.Service, util.Degub)
	if err != nil {
		log.Fatal("cannot create logger")
	}
	defer closeFunc()

	querier, err := db.NewQuerier(config, logger)
	if err != nil {
		logger.Criticalf(context.Background(), "cannot connect to db:", err)
	}
	defer querier.DB.Close()

	matcher := util.NewMatcher()

	// DI to server
	server, err := api.NewServer(config, querier, matcher, logger)
	if err != nil {
		logger.Criticalf(context.Background(), "cannot create server:", err)
	}

	if config.IsCGI {
		server.CGI()
	} else {
		err = server.Start()
		if err != nil {
			logger.Criticalf(context.Background(), "cannot start server:", err)
		}
	}
}
