package database

import (
	"fmt"
	"os"
	"path/filepath"
)

// Delete a user photo.
// Return a payload feedback message and error
func (db *appdbimpl) DeletePhoto(idphoto string) (string, error) {

	abs, errPath := filepath.Abs(".")
	if errPath != nil {
		// http status 500
		return "error processing abs path", errPath
	}

	filepath := filepath.Join(abs, "storage", idphoto+".png")

	err := os.Remove(filepath)
	if err != nil {
		return "The server cannot delete the file", err
	}

	sqlStmt := fmt.Sprintf("DELETE FROM Photo WHERE Id_photo = '%s';", idphoto)

	_, errQuery := db.c.Exec(sqlStmt)

	if errQuery != nil {
		return "error execution query", fmt.Errorf("error execution DeletePhoto query: %w", errQuery)
	}

	return "Image deleted, everything ok", nil
}
