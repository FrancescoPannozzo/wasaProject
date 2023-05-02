package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CheckBan(loggedUser string, targetUser string) bool {
	var banned string
	rows := db.c.QueryRow("SELECT Banned FROM Ban WHERE Banner =? AND Banned =?;", targetUser, loggedUser).Scan(&banned)

	if errors.Is(rows, sql.ErrNoRows) {
		return false
	}

	return true
}
