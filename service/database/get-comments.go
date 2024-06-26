package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

func (db *appdbimpl) GetComments(loggedUser string, photoID string) ([]utilities.Comment, error) {
	var comments []utilities.Comment

	rows, err := db.c.Query("SELECT Id_comment, User, Content FROM Comment WHERE Photo =?;", photoID)
	if err != nil {
		// http status 500
		return nil, fmt.Errorf("error while getting the comment list: %w", err)
	}

	var comment utilities.Comment
	for rows.Next() {

		var commentId, user, content string
		errScan := rows.Scan(&commentId, &user, &content)
		if errScan != nil {
			return nil, fmt.Errorf("error while scanning the comment list: %w", errScan)
		}
		if db.CheckBan(loggedUser, user) {
			continue
		}
		comment.CommentId = commentId
		comment.Name = user
		comment.Content = content
		comments = append(comments, comment)
	}
	errScan := rows.Err()
	if errScan != nil {
		return nil, fmt.Errorf("Error while scanning for GetComments operation: %w", errScan)
	}

	// http status 200
	return comments, nil
}
