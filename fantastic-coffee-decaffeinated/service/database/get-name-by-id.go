package database

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

// Get an username by passing the userId associated from the DB,
func (db *appdbimpl) GetNameByID(userId string) (string, error) {

	var (
		userID   string
		nickname string
	)

	logrus.Infoln("Getting the nickaname for userId ", userId)

	rows := db.c.QueryRow("SELECT Nickname FROM User WHERE Id_user=?", userId).Scan(&nickname)

	if errors.Is(rows, sql.ErrNoRows) {
		logrus.Printf("UserId %s not in the db", userId)
		//errUser := fmt.Errorf("error execution query: %w", rows)
		return "No user found", rows
	}

	logrus.Printf("UserId: %s found! nickname is: %s\n", userID, nickname)
	return nickname, nil
}
