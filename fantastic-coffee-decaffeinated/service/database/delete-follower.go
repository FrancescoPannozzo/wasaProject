package database

import (
	"fmt"
	"net/http"
)

// Delete a follow
func (db *appdbimpl) DeleteFollowed(follower string, followed string) (string, error, int) {

	_, err := db.c.Exec("DELETE FROM Follow WHERE Follower = ? AND Followed = ?", follower, followed)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "unfollow done, ok", err, http.StatusOK

}
