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
func (db *appdbimpl) DeletePhoto(name string, idphoto string) (string, error, int) {

	abs, errPath := filepath.Abs(".")

	fmt.Printf("abs for . is:%s\nerr:%v\n", abs, errPath)

	filepath := filepath.Join(abs, "storage", idphoto+".png")

	fmt.Println(filepath)

	checkFileExists(filepath)

	e := os.Remove(filepath)
	if e != nil {
		return "The server cannot delete the file", e, http.StatusInternalServerError
	}

	sqlStmt := fmt.Sprintf("DELETE FROM Photo WHERE Id_photo = '%s';", idphoto)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "error execution query", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "Image deleted, everything ok", err, http.StatusOK

}
