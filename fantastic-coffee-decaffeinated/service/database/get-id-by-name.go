package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Get an username identifier from the DB,
func (db *appdbimpl) GetIdByName(name string) (string, error) {

	var (
		userID   string
		username string
	)

	//fmt.Println("now in CheckUser, name value is:", name)

	logrus.Infoln("Getting the user ID..")

	rows := db.c.QueryRow("SELECT Id_user, Nickname FROM User WHERE Nickname=?", name).Scan(&userID, &username)
	if errors.Is(rows, sql.ErrNoRows) {
		logrus.Println("User not in the db")
		errUser := fmt.Errorf("error execution query: %w", rows)
		return "", errUser
	}

	logrus.Printf("User: %s found! user ID is: %v\n", username, userID)
	return userID, nil
}
