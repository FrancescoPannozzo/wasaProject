package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
)

// Insert the user in the DB,
func (db *appdbimpl) LikePhoto(username string, idphoto string) (string, error) {

	var user string
	rows := db.c.QueryRow("SELECT User FROM Like WHERE User=? AND Photo=?", username, idphoto).Scan(&user)
	if !errors.Is(rows, sql.ErrNoRows) {
		return "warning: the user already like the target user photo", &utilities.DbBadRequest{}
	}

	_, err := db.c.Exec("INSERT INTO Like (User, Photo) VALUES(?, ?);", username, idphoto)
	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", err
	}
	//201
	return "like added, ok", nil
}
