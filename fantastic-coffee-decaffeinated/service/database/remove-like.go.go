package database

import (
	"fmt"
	"net/http"
)

// Delete a follow
func (db *appdbimpl) RemoveLike(username string, idphoto string) (string, error, int) {
	_, err := db.c.Exec("DELETE FROM Like WHERE User = ? AND Photo = ?", username, idphoto)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "like removed done, ok", err, http.StatusOK

}
