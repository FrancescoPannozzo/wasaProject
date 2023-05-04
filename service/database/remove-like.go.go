package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Delete a like
func (db *appdbimpl) RemoveLike(loggedUser string, idphoto string) (string, error) {
	var username string

	rows := db.c.QueryRow("SELECT User FROM Like WHERE User=? AND Photo=?", loggedUser, idphoto).Scan(&username)

	if errors.Is(rows, sql.ErrNoRows) {
		// 404 like not found
		return "like not found", &utilities.DbBadRequestError{}
	}

	_, err := db.c.Exec("DELETE FROM Like WHERE User = ? AND Photo = ?", username, idphoto)

	if err != nil {
		// 500 Internal server error
		return utilities.ErrorExecutionQuery, fmt.Errorf("error execution query in Like: %w", err)
	}

	// http status 200
	return "like removed done, ok", nil

}
