package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Ban the provided user. Returns a feedback string and nil if successful,
// a feedback string and error != nil otherwise.
func (db *appdbimpl) BanUser(banner string, banned string) (string, error) {
	var user string
	rows := db.c.QueryRow("SELECT Banner FROM Ban WHERE Banner=? AND Banned=?", banner, banned).Scan(&user)
	if !errors.Is(rows, sql.ErrNoRows) {
		return "Warning, the user is already banned from the logged user", &utilities.DbBadRequestError{}
	}

	sqlStmt := fmt.Sprintf("INSERT INTO Ban (Banner, Banned) VALUES('%s','%s');", banner, banned)
	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// http status 500
		return "cannot insert the ban into the DB", fmt.Errorf("error execution query: %w", err)
	}
	// http status 201
	return "Banned user inserted in the DB", nil
}
