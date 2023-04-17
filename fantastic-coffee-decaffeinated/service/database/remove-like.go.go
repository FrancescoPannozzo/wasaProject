package database

import (
	"fmt"
)

// Delete a follow
func (db *appdbimpl) RemoveLike(username string, idphoto string) (string, error) {
	_, err := db.c.Exec("DELETE FROM Like WHERE User = ? AND Photo = ?", username, idphoto)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "like removed done, ok", nil

}
