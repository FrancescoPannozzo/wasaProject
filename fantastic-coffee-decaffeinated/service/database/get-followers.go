package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Return the logged user followers
func (db *appdbimpl) GetFollowers(loggedUser string) ([]utilities.Username, error) {

	var followers []utilities.Username

	rows, err := db.c.Query("SELECT Follower FROM Follow WHERE Followed=?;", loggedUser)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var username utilities.Username
	for rows.Next() {
		rows.Scan(&username.Name)
		followers = append(followers, username)
	}

	return followers, nil
}
