package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// @todo: e' l'handler che decide che tipo di http status code mandare, nelle data abse functions ritornare solo la roba richiesta ed err
func (db *appdbimpl) GetUsernames(targetUser string) ([]utilities.Username, error) {

	var usernames []utilities.Username

	querytStmt := fmt.Sprintf("SELECT Nickname FROM User WHERE Nickname LIKE '%%%s%%';", targetUser)

	rows, err := db.c.Query(querytStmt)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var username utilities.Username
	for rows.Next() {
		rows.Scan(&username.Name)
		usernames = append(usernames, username)
	}

	return usernames, nil
}
