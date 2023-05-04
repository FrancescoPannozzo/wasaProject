package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

func (db *appdbimpl) UnbanUser(banner string, banned string) (string, error) {
	var bannerUser string
	rows := db.c.QueryRow("SELECT Banner FROM Ban WHERE Banned=? AND Banner=?;", banned, banner).Scan(&bannerUser)
	if errors.Is(rows, sql.ErrNoRows) {
		// 404 like not found
		return "Ban not found", &utilities.DbBadRequestError{}
	}

	sqlStmt := fmt.Sprintf("DELETE FROM Ban WHERE Banner='%s' AND Banned='%s';", banner, banned)
	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// http status 500
		return "error execution query", fmt.Errorf("error execution query: %w", err)
	}

	// http status 200
	return "User unbanned, DB updated", nil
}
