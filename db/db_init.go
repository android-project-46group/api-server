package db

import (
	"context"
	"database/sql"

	"github.com/android-project-46group/api-server/util"
	_ "github.com/lib/pq"
)

// Implementation of Querier interface
type SqlQuerier struct {
	DB  *sql.DB
	ctx context.Context
}

func NewQuerier(config util.Config) (*SqlQuerier, error) {

	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, err
	}

	querier := &SqlQuerier{
		DB:  con,
		ctx: context.Background(),
	}

	return querier, nil
}
