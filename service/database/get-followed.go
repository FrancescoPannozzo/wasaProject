package database

import (
	"fmt"
)

// Return the logged user followers
func (db *appdbimpl) GetFollowed(loggedUser string) ([]string, error) {

	var followed []string

	rows, err := db.c.Query("SELECT Followed FROM Follow WHERE Follower=?;", loggedUser)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var username string
	for rows.Next() {
		errScan := rows.Scan(&username)
		if errScan != nil {
			return nil, fmt.Errorf("error while scanning Follow: %w", errScan)
		}
		followed = append(followed, username)
	}

	return followed, nil
}
