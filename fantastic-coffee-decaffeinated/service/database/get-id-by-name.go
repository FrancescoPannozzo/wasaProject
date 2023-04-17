package database

import (
	"database/sql"
	"errors"
)

// Get an username identifier from the DB. @name the username
// If the username is present in the DB it returns the userId, nil
// It returns "", error otherwise.
func (db *appdbimpl) GetIdByName(name string) (string, error) {

	var (
		userID   string
		username string
	)

	rows := db.c.QueryRow("SELECT Id_user, Nickname FROM User WHERE Nickname=?", name).Scan(&userID, &username)
	if errors.Is(rows, sql.ErrNoRows) {
		return "", rows
	}

	return userID, nil
}
