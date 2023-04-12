package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (db *appdbimpl) GetUsernames(targetUser string) ([]utilities.Username, error, int) {

	var usernames []utilities.Username

	querytStmt := fmt.Sprintf("SELECT Nickname FROM User WHERE Nickname LIKE '%%%s%%';", targetUser)
	logrus.Println(querytStmt)

	rows, err := db.c.Query(querytStmt)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	var username utilities.Username
	for rows.Next() {
		rows.Scan(&username.Name)
		usernames = append(usernames, username)
		fmt.Println(username.Name)
	}

	return usernames, nil, http.StatusOK

}
