package database

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

// Get an username identifier from the DB,
// If the username is present in the DB it returns the userId, nil
// if not it returns "", error
func (db *appdbimpl) GetIdByName(name string) (string, error) {

	var (
		userID   string
		username string
	)

	//fmt.Println("now in CheckUser, name value is:", name)

	logrus.Infoln("Getting the user ID for username ", name)

	rows := db.c.QueryRow("SELECT Id_user, Nickname FROM User WHERE Nickname=?", name).Scan(&userID, &username)

	if errors.Is(rows, sql.ErrNoRows) {
		logrus.Printf("User %s not in the db", name)
		//errUser := fmt.Errorf("error execution query: %w", rows)
		return "", rows
	}

	logrus.Printf("User: %s found! user ID is: %v\n", username, userID)
	return userID, nil
}
