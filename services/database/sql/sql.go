package sql

import (
	"database/sql"
	"fmt"
)

// Connect func
func Connect(dlct, user, pass, name, host, port string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		user, pass, name, host, port)
	return sql.Open(dlct, connStr)
}
