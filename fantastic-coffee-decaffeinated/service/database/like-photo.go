package database

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

// Insert the user in the DB,
func (db *appdbimpl) LikePhoto(username string, idphoto string) (string, error, int) {
	// check if idphoto is valid ---------

	// @todo: potevo fare funzioncina check id photo
	var idPhoto string
	rows := db.c.QueryRow("SELECT DISTINCT Id_photo FROM Photo WHERE Id_photo=?", idphoto).Scan(&idPhoto)

	if errors.Is(rows, sql.ErrNoRows) {
		//errUser := fmt.Errorf("error execution query: %w", rows)
		return "invalid idphoto", rows, http.StatusBadRequest
	}

	_, err := db.c.Exec("INSERT INTO Like (User, Photo) VALUES(?, ?);", username, idphoto)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "like added, ok", nil, http.StatusCreated

}
