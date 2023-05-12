package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Insert the user in the DB,
// return the userID and nil (status created) if successfull
// If there is an internal error it return empty userID and error != nil (http status code 500)
func (db *appdbimpl) InsertUser(name string) (string, error) {
	// create user id
	userID := utilities.GenerateUserID()

	logrus.Infof("User id created: %s, inserting the new user in the db..\n", userID)

	sqlStmt := fmt.Sprintf("INSERT INTO User (Id_user, Nickname) VALUES('%s','%s');", userID, name)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "", fmt.Errorf("error execution query: %w", err)
	}
	logrus.Infof("..done!")
	// 201 Created
	return userID, nil
}
