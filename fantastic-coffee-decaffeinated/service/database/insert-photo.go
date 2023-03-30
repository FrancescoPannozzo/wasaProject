package database

import (
	"fmt"
	"net/http"
	"time"
)

// Insert the user photo in the DB,
func (db *appdbimpl) InsertPhoto(name string, idphoto string) (string, error, int) {

	userID, _ := DBcon.GetIdByName(name)

	now := time.Now()

	date := now.Format("2006-01-02")

	time := now.Format("15:04:05")

	sqlStmt := fmt.Sprintf("INSERT INTO Photo (Id_photo, User, Date, Time, LocalPath) VALUES('%s', '%s','%s','%s', '%s');", idphoto, userID, date, time, idphoto+".png")

	fmt.Println(sqlStmt)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "error execution query", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	// 201 Created
	return "Photo inserted in the DB", nil, http.StatusCreated
}
