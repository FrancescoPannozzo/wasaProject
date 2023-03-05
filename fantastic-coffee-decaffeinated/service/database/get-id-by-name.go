package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (db *appdbimpl) GetIdByName(name string) (string, error) {

	var (
		userID   string
		username string
	)

	//fmt.Println("now in CheckUser, name value is:", name)

	logrus.Infoln("now in GetIdByName(), name value is:", name)

	rows := db.c.QueryRow("SELECT Id_user, Nickname FROM User WHERE Nickname=?", name).Scan(&userID, &username)
	if errors.Is(rows, sql.ErrNoRows) {
		logrus.Infoln("User not in the db")
		errUser := fmt.Errorf("error execution query: %w", rows)
		return "", errUser
	}

	logrus.Infoln("User: %s already in the db\n", username)
	return userID, nil
}
