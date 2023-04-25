package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

func (db *appdbimpl) UnbanUser(banner string, banned string) (string, error) {
	var bannedUser string
	rows := db.c.QueryRow("SELECT Banner FROM Ban WHERE Banned=?", banned).Scan(&bannedUser)
	if errors.Is(rows, sql.ErrNoRows) {
		// 404 like not found
		return "Ban not found", &utilities.DbBadRequest{}
	}

	sqlStmt := fmt.Sprintf("DELETE FROM Ban WHERE Banner='%s' AND Banned='%s';", banner, banned)
	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		//500
		return "error execution query", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "User unbanned, DB updated", err
}
