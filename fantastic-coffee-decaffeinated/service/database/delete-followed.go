package database

import (
	"fmt"
)

// Delete a followed user. Possible http status code: 200, 500
func (db *appdbimpl) DeleteFollowed(follower string, followed string) (string, error) {

	_, err := db.c.Exec("DELETE FROM Follow WHERE Follower = ? AND Followed = ?", follower, followed)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "unfollow done, ok", err

}
