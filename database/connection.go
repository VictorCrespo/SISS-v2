package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	db   *sql.DB
)

func GetConnectionPool() (*sql.DB, error) {

	var err error

	once.Do(func() {

		username, ok := os.LookupEnv("DB_USERNAME")

		if !ok {
			err = fmt.Errorf("DB_USERNAME not defined")
			return
		}

		password, ok := os.LookupEnv("DB_PASSWORD")

		if !ok {
			err = fmt.Errorf("DB_PASSWORD not defined")
			return
		}

		hostname, ok := os.LookupEnv("DB_HOSTNAME")

		if !ok {
			err = fmt.Errorf("DB_HOSTNAME not defined")
			return
		}

		port, ok := os.LookupEnv("DB_PORT")

		if !ok {
			err = fmt.Errorf("DB_PORT not defined")
			return
		}

		dbname, ok := os.LookupEnv("DB_NAME")

		if !ok {
			err = fmt.Errorf("DB_NAME not defined")
			return
		}

		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, password, hostname, port, dbname))

		if err != nil {
			err = fmt.Errorf("failed to connect to database: %v", err)
			return
		}

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(3 * time.Minute)
	})

	return db, err
}
