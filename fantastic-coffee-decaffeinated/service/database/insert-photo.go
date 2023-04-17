package database

import (
	"fmt"
	"time"
)

// Insert the user photo in the DB, returns a feedback string and nil if succesfull,
// a feedback string and an error otherwise.
func (db *appdbimpl) InsertPhoto(name string, idphoto string) (string, error) {
	now := time.Now()
	date := now.Format("2006-01-02")
	time := now.Format("15:04:05")

	sqlStmt := fmt.Sprintf("INSERT INTO Photo (Id_photo, User, Date, Time, LocalPath) VALUES('%s', '%s','%s','%s', '%s');", idphoto, name, date, time, idphoto+".png")
	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "error execution query", fmt.Errorf("error execution query: %w", err)
	}

	// 201 Created
	return "Photo inserted in the DB", nil
}
