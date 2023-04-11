package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"
)

func (db *appdbimpl) GetComments(loggedUser string, photoID string) ([]utilities.Comment, error, int) {
	//var idphoto string

	var comments []utilities.Comment

	//Id_comment|User|Photo|Content

	rows, err := db.c.Query("SELECT User, Content FROM Comment WHERE Photo =?;", photoID)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
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

	return comments, nil, http.StatusOK

}
