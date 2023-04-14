package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

func (db *appdbimpl) GetComments(loggedUser string, photoID string) ([]utilities.Comment, error) {
	//var idphoto string

	var comments []utilities.Comment

	//Id_comment|User|Photo|Content

	rows, err := db.c.Query("SELECT User, Content FROM Comment WHERE Photo =?;", photoID)
	if err != nil {
		//500
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var comment utilities.Comment
	for rows.Next() {

		var user, content string
		rows.Scan(&user, &content)
		if db.CheckBan(loggedUser, user) {
			continue
		}
		comment.Name = user
		comment.Content = content
		comments = append(comments, comment)

		fmt.Println(comments)
	}

	//200
	return comments, nil

}
