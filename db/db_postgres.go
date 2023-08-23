package db

import (
	"database/sql"
	"fmt"

	"github.com/kuma-coffee/go-rest-api/config"
)

func CreateConn() *sql.DB {
	conf := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s"+"password=%s dbname=%s sslmode=disable", conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
