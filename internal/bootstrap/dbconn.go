// Package bootstrap stores basic common entities such as config and database connection
package bootstrap

import (
	"fmt"

	"github.com/gocraft/dbr"
)

const maxOpenConns = 10

// NewDBConn creates a new database connection (connection pool) instance.
func NewDBConn(
	scheme,
	username,
	password,
	name,
	host,
	port string,
) (*dbr.Connection, error) {
	// opening new connection; it's NOT necessary to close it
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		username,
		password,
		name,
		host,
		port,
	)

	p, err := dbr.Open(scheme, dsn, nil)
	if err != nil {
		return nil, err
	}

	// verifying that it's alive at the beginning
	err = p.Ping()
	if err != nil {
		return nil, err
	}

	// a maximum of 10 concurrent connections; might be changed later if needed
	p.SetMaxOpenConns(maxOpenConns)

	return p, nil
}
