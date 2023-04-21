package database

import (
	"fmt"
)

// Delete a follow
func (db *appdbimpl) RemoveComment(idcomment string) (string, error) {
	_, err := db.c.Exec("DELETE FROM Comment WHERE Id_comment = ?;", idcomment)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "comment removed, ok", nil

}
