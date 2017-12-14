package config

import (
	"database/sql"
	"log"
	"os"

	db "github.com/tsrnd/go-clean-arch/services/database/sql"
)

// DB func
func DB() *sql.DB {
	dbDlct := os.Getenv("DATABASE_DLCT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	db, err := db.Connect(dbDlct, dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
