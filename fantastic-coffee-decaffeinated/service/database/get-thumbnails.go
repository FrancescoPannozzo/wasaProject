package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Get user thumbnails objects.
func (db *appdbimpl) GetThumbnails(username string) ([]utilities.Thumbnail, error) {
	//var idphoto string

	var thumbnails []utilities.Thumbnail

	rows, err := db.c.Query("SELECT User, Id_photo, Date, Time FROM Photo WHERE User=? ORDER BY Date DESC, Time Desc;", username)
	if err != nil {
		//500
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var thumbnail utilities.Thumbnail
	for rows.Next() {

		var date, time string
		rows.Scan(&thumbnail.Username, &thumbnail.PhotoId, &date, &time)
		thumbnail.DateTime = fmt.Sprintf("%sT%s", date, time)
		rows := db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE Photo = ?;", thumbnail.PhotoId).Scan(&thumbnail.LikesNumber)
		if errors.Is(rows, sql.ErrNoRows) {
			//500
			return nil, fmt.Errorf("error execution query: %w", rows)
		}
		rows = db.c.QueryRow("SELECT COUNT(*) FROM Comment WHERE Photo = ?;", thumbnail.PhotoId).Scan(&thumbnail.CommentsNumber)
		if errors.Is(rows, sql.ErrNoRows) {
			//500
			return nil, fmt.Errorf("error execution query: %w", rows)
		}
		thumbnail.PhotoURL = utilities.CreatePhotoURL(thumbnail.PhotoId)
		thumbnails = append(thumbnails, thumbnail)
	}

	//200
	return thumbnails, nil

}
