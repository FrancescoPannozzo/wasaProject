package database

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func checkFileExists(filepath string) {
	_, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file not exists")
	} else {
		fmt.Println("file exists")
	}
}

// Delete a user photo.
// Return a payload feedback message and error
func (db *appdbimpl) DeletePhoto(idphoto string) (string, error) {

	abs, errPath := filepath.Abs(".")
	if errPath != nil {
		//500
		return "error processing abs path", errPath
	}

	filepath := filepath.Join(abs, "storage", idphoto+".png")

	checkFileExists(filepath)

	err := os.Remove(filepath)
	if err != nil {
		return "The server cannot delete the file", err
	}

	sqlStmt := fmt.Sprintf("DELETE FROM Photo WHERE Id_photo = '%s';", idphoto)

	_, errQuery := db.c.Exec(sqlStmt)

	if errQuery != nil {
		return "error execution query", fmt.Errorf("error execution query: %w", errQuery)
	}

	return "Image deleted, everything ok", nil

}
