package db

import (
	"context"
	"database/sql"
	"fmt"
	
	_ "github.com/lib/pq"
)


var (
	// TODO: get from env
	// DbPass = os.Getenv("DBPASSWORD")
	DbName = "sakamichi"
	DbPass = "sakamichi"
	
	ConnectionString = 
		fmt.Sprintf("host=localhost port=5432 user=ubuntu password=%s dbname=%s sslmode=disable", DbPass, DbName)

	DB  *sql.DB
	Ctx context.Context
)

func DbInit() (*sql.DB, error) {
	con, err := Connect()
	Ctx = context.Background()
	return con, err
}

func Connect() (*sql.DB, error){
	connection, err := sql.Open("postgres", ConnectionString)
	DB = connection
	
	return connection, err
}
