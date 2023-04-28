package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Return the logged user followers
func (db *appdbimpl) GetFollowed(loggedUser string) ([]utilities.Username, error) {

	var followed []utilities.Username

	rows, err := db.c.Query("SELECT Followed FROM Follow WHERE Follower=?;", loggedUser)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var username utilities.Username
	for rows.Next() {
		rows.Scan(&username.Name)
		followed = append(followed, username)
	}

	return followed, nil
}
