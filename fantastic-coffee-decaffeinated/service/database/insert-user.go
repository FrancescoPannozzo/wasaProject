package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Insert the user in the DB,
// return the userID, error = nil and the the http.StatusCreated
// If there is an internal error it return empty userID, error != nil and http.StatusInternalServerError
func (db *appdbimpl) InsertUser(name string) (string, error, int) {
	//create user id
	userID := utilities.GenerateUserID(name)

	logrus.Infof("User id created =%s, inserting the new user in the db..\n", userID)

	sqlStmt := fmt.Sprintf("INSERT INTO User (Id_user, Nickname) VALUES('%s','%s');", userID, name)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	// 201 Created
	return userID, nil, http.StatusCreated
}
