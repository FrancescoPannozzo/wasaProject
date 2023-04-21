package database

import (
	"database/sql"
	"errors"
)

// Get the username of the provided commentId.
// Return "" and sql.ErrNoRows error if not found.
// Return the username and nil otherwise.
func (db *appdbimpl) GetNameFromCommentId(commentId string) (string, error) {
	var username string

	rows := db.c.QueryRow("SELECT User FROM Comment WHERE Id_comment=?", commentId).Scan(&username)

	if errors.Is(rows, sql.ErrNoRows) {
		return "", rows
	}

	return username, nil
}
