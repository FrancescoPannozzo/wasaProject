package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Delete a followed user. Possible http status code: 200, 500
func (db *appdbimpl) DeleteFollowed(follower string, followed string) (string, error) {
	var user string
	rows := db.c.QueryRow("SELECT Follower FROM Follow WHERE Follower=? AND Followed=?", follower, followed).Scan(&user)
	if errors.Is(rows, sql.ErrNoRows) {
		return "warning: user not found for the unfollow operation", &utilities.DbBadRequest{}
	}

	_, err := db.c.Exec("DELETE FROM Follow WHERE Follower = ? AND Followed = ?", follower, followed)
	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "unfollow done, ok", err

}
