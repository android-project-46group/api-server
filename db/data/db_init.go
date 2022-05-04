package data

import (
	"database/sql"
	"fmt"
	"context"
	
	_ "github.com/lib/pq"
)


var (
	// TODO: get from env
	// DbPass = os.Getenv("DBPASSWORD")
	DbName = "sakamichi"
	DbPass = "sakamichi"
	
	UserInfo   = "host=localhost port=5432 user=ubuntu password=" + DbPass + " dbname=" + DbName + " sslmode=disable"

	DB  *sql.DB
	Ctx context.Context
)

func DbInit() (*sql.DB, error) {
	con, err := Connect()
	Ctx = context.Background()
	return con, err
}

func Connect() (*sql.DB, error){
	connection, err := sql.Open("postgres", UserInfo)
	DB = connection
	fmt.Println(DB)
	return connection, err
}
