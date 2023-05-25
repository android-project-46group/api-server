package main

import (
	"context"
	"log"
	"os"

	"github.com/android-project-46group/api-server/api"
	"github.com/android-project-46group/api-server/db"
	"github.com/android-project-46group/api-server/repository/grpc"
	"github.com/android-project-46group/api-server/util"
	"github.com/opentracing/opentracing-go"
)

func main() {
	tr, closer, err := util.NewJaegerTracer()
	if err != nil {
		log.Fatal("cannot initialize jaeger tracer: ", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tr)

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	logger, closeFunc, err := util.NewZapLogger(config.LogPath, config.Host, config.Service, util.Degub)
	// logger, closeFunc, err := util.NewStandardLogger(config.Host, config.Service, os.Stdout)
	// logger.SetLevel(util.Degub)
	if err != nil {
		log.Fatal("cannot create logger: ", err)
	}
	defer closeFunc()

	querier, err := db.NewQuerier(config, logger)
	if err != nil {
		logger.Criticalf(context.Background(), "cannot connect to db: %v", err)
		os.Exit(1)
	}
	defer querier.DB.Close()

	matcher := util.NewMatcher()

	downloadClient, grpcCloserFunc, err := grpc.NewGRPCClient(config.GRPCURL)
	if err != nil {
		logger.Criticalf(context.Background(), "failed to NewGRPCClient: %v", err)
		os.Exit(1)
	}
	defer grpcCloserFunc()

	grpcClient := grpc.NewClient(logger, config, downloadClient)

	// DI to server
	server, err := api.NewServer(config, querier, matcher, logger, grpcClient)
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
