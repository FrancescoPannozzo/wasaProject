package database

import (
	"fmt"
)

// Get a list of usernames
func (db *appdbimpl) GetUsernames(targetUser string) ([]string, error) {
	var usernames []string

	querytStmt := fmt.Sprintf("SELECT Nickname FROM User WHERE Nickname LIKE '%%%s%%';", targetUser)

	rows, err := db.c.Query(querytStmt)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var username string
	for rows.Next() {
		rows.Scan(&username)
		usernames = append(usernames, username)
	}

	return usernames, nil
}
