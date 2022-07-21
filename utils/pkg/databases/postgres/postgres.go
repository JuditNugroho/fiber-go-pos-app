package postgres

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbPG *sqlx.DB

func OpenConnection() error {
	var err error

	// Initialize variable
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)

	if err != nil {
		return errors.New("port must be number")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	dbPG, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		return err
	}

	dbPG.SetMaxIdleConns(10)
	dbPG.SetConnMaxLifetime(5 * time.Minute)
	dbPG.SetMaxOpenConns(50)

	return nil
}

func GetPgConn() *sqlx.DB {
	return dbPG
}
