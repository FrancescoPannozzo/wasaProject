package database

import (
	"fmt"
	"net/http"
)

// Insert the user in the DB,
func (db *appdbimpl) LikePhoto(username string, idphoto string) (string, error, int) {
	/*
		idphotoConv, errconv := strconv.Atoi(idphoto)

		if errconv != nil {
			logrus.Errorln("error while processing the photo request")
			return "error while processing the photo request", errconv, http.StatusInternalServerError
		}
	*/
	_, err := db.c.Exec("INSERT INTO Like (User, Photo) VALUES(?, ?);", username, idphoto)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "like added, ok", err, http.StatusCreated

}
