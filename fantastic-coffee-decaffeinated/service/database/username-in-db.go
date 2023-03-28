package database

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (db *appdbimpl) UsernameInDB(name string) bool {
	var username string
	rows := db.c.QueryRow("SELECT Nickname FROM User WHERE Nickname = ?;", name).Scan(&username)

	if errors.Is(rows, sql.ErrNoRows) {
		logrus.Printf("User %s not in the db", name)
		return false
	}
	return true
}
