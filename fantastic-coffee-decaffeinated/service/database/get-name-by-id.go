package database

import (
	"database/sql"
	"errors"
)

// Get an username by passing the userId associated from the DB.
// return the username and nil if sucessful, a feedback message anthe ErrNoRows otherwise
func (db *appdbimpl) GetNameByID(userId string) (string, error) {

	var nickname string

	rows := db.c.QueryRow("SELECT Nickname FROM User WHERE Id_user=?", userId).Scan(&nickname)
	if errors.Is(rows, sql.ErrNoRows) {
		return "", rows
	}

	return nickname, nil
}
