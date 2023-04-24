package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

func (db *appdbimpl) GetComments(loggedUser string, photoID string) ([]utilities.Comment, error) {
	var photo string

	row := db.c.QueryRow("SELECT Photo FROM Comment WHERE Photo =?;", photoID).Scan(&photo)
	if errors.Is(row, sql.ErrNoRows) {
		//404 photo not found in Comment DB table, comments not found
		return nil, row
	}

	var comments []utilities.Comment

	rows, err := db.c.Query("SELECT Id_comment, User, Content FROM Comment WHERE Photo =?;", photoID)
	if err != nil {
		//500
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var comment utilities.Comment
	for rows.Next() {

		var commentId, user, content string
		rows.Scan(&commentId, &user, &content)
		if db.CheckBan(loggedUser, user) {
			continue
		}
		comment.CommentId = commentId
		comment.Name = user
		comment.Content = content
		comments = append(comments, comment)
	}

	//200
	return comments, nil
}
