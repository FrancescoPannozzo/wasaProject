package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) UsernameInDB(name string) bool {
	var username string
	rows := db.c.QueryRow("SELECT Nickname FROM User WHERE Nickname = ?;", name).Scan(&username)

	if errors.Is(rows, sql.ErrNoRows) {
		return false
	}
	return true
}
