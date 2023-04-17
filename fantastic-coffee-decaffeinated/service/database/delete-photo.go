package database

import (
	"errors"
	"fmt"
	"net/http"
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

// Delete an user photo
// return a payload feedback message, an error and a http status code
func (db *appdbimpl) DeletePhoto(idphoto string) (string, error, int) {

	abs, errPath := filepath.Abs(".")
	if errPath != nil {
		return "error processing abs path", errPath, http.StatusInternalServerError
	}

	filepath := filepath.Join(abs, "storage", idphoto+".png")

	checkFileExists(filepath)

	e := os.Remove(filepath)
	if e != nil {
		return "The server cannot delete the file", e, http.StatusInternalServerError
	}

	sqlStmt := fmt.Sprintf("DELETE FROM Photo WHERE Id_photo = '%s';", idphoto)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return "error execution query", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "Image deleted, everything ok", err, http.StatusOK

}
