package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Delete a follow
func (db *appdbimpl) RemoveComment(idcomment string) (string, error) {
	var username string

	rows := db.c.QueryRow("SELECT User FROM Comment WHERE Id_comment=?", idcomment).Scan(&username)

	if errors.Is(rows, sql.ErrNoRows) {
		// 404 comment not found
		return "comment not found", &utilities.DbBadRequest{}
	}

	_, err := db.c.Exec("DELETE FROM Comment WHERE Id_comment = ?;", idcomment)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "comment removed, ok", nil

}
