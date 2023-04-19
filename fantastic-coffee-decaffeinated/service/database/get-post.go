package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Get user thumbnails objects.
func (db *appdbimpl) GetPost(loggedUser string, photoId string) (utilities.Post, error) {
	//var idphoto string

	var post utilities.Post
	var date, time string

	rows := db.c.QueryRow("SELECT Date, Time, Photo_url FROM Photo WHERE Id_photo = ?;", photoId).Scan(&date, &time, &post.PhotoURL)
	if errors.Is(rows, sql.ErrNoRows) {
		//500
		return post, fmt.Errorf("error execution query: %w", rows)
	}

	post.DateTime = fmt.Sprintf("%sT%s", date, time)

	rows = db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE Photo = ?;", photoId).Scan(&post.LikesNumber)
	if errors.Is(rows, sql.ErrNoRows) {
		//500
		return post, fmt.Errorf("error execution query: %w", rows)
	}

	comments, err := db.GetComments(loggedUser, photoId)
	if err != nil {
		return post, err
	}

	post.Comments = comments

	//200
	return post, nil

}
