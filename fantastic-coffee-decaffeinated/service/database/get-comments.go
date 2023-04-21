package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (db *appdbimpl) GetComments(loggedUser string, photoID string) ([]utilities.Comment, error) {
	logrus.Infoln("Getting the comments..")
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
	logrus.Infoln("Done!")
	//200
	return comments, nil
}
