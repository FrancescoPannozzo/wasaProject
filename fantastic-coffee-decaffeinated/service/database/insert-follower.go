package database

import (
	"fmt"
	"net/http"
)

// Insert the user in the DB,
func (db *appdbimpl) InsertFollower(follower string, followed string) (string, error, int) {
	sqlStmt := fmt.Sprintf("INSERT INTO Follow (Follower, Followed) VALUES('%s','%s');", follower, followed)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "follow added, ok", err, http.StatusCreated

}
