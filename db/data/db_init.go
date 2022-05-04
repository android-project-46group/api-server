package data

import (
	"context"
	"database/sql"
	"log"

	"github.com/android-project-46group/api-server/util"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	Ctx context.Context
)

func DbInit() (*sql.DB, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	Ctx = context.Background()
	DB, err = sql.Open(config.DBDriver, config.DBSource)
	return DB, err
}
