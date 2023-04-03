package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetNameFromPhotoId(photoId string) (string, error) {
	var username string

	rows := db.c.QueryRow("SELECT User FROM Photo WHERE Id_photo=?", photoId).Scan(&username)

	if errors.Is(rows, sql.ErrNoRows) {
		return "", rows
	}

	return username, nil
}
