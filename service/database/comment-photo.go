package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Insert the comment on the photoID provided in the DB.
// Return a feedback message and nil if successfull.
// Return a feedback message and an error excetution query otherwise.
func (db *appdbimpl) CommentPhoto(username string, idphoto string, comment string) (string, error) {
	_, err := db.c.Exec("INSERT INTO Comment (User, Photo, Content) VALUES(?, ?, ?);", username, idphoto, comment)

	if err != nil {
		// 500 Internal server error
		return utilities.ErrorExecutionQuery, fmt.Errorf("error execution query: %w", err)
	}
	// 200 status created
	return "comment added, ok", nil

}
