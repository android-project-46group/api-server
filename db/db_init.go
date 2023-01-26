package db

import (
	"database/sql"
	"fmt"

	"github.com/android-project-46group/api-server/util"
	_ "github.com/lib/pq"
)

// Implementation of Querier interface
type SqlQuerier struct {
	DB *sql.DB
}

func NewQuerier(config util.Config) (*SqlQuerier, error) {

	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("NewQuerier: %w", err)
	}

	querier := &SqlQuerier{
		DB: con,
	}

	return querier, nil
}
