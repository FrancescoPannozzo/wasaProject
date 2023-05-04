package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Get user thumbnails objects.
func (db *appdbimpl) GetPost(loggedUser string, photoId string) (utilities.Post, error) {
	var post utilities.Post
	var date, time string

	rows := db.c.QueryRow("SELECT User, Date, Time, Photo_url FROM Photo WHERE Id_photo = ?;", photoId).Scan(&post.Username, &date, &time, &post.PhotoURL)

	if errors.Is(rows, sql.ErrNoRows) {
		// http status 500
		return post, fmt.Errorf("Error execution query: %w", rows)
	}

	post.DateTime = fmt.Sprintf("%sT%s", date, time)

	rows = db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE Photo = ?;", photoId).Scan(&post.LikesNumber)
	if errors.Is(rows, sql.ErrNoRows) {
		// http status 500
		return post, fmt.Errorf("error execution query: %w", rows)
	}

	comments, err := db.GetComments(loggedUser, photoId)
	if err != nil {
		return post, fmt.Errorf("error getting comments: %w", err)
	}

	post.Comments = comments

	// http status 200
	return post, nil

}
