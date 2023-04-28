package database

import (
	"fmt"
)

// Return the logged user followers
func (db *appdbimpl) GetFollowers(loggedUser string) ([]string, error) {

	var followers []string

	rows, err := db.c.Query("SELECT Follower FROM Follow WHERE Followed=?;", loggedUser)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var username string
	for rows.Next() {
		rows.Scan(&username)
		followers = append(followers, username)
	}

	return followers, nil
}
