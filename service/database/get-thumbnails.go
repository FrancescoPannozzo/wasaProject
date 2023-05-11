package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Get user thumbnails objects.
func (db *appdbimpl) GetThumbnails(username string) ([]utilities.Thumbnail, error) {
	var thumbnails []utilities.Thumbnail

	rows, err := db.c.Query("SELECT User, Id_photo, Date, Time FROM Photo WHERE User=? ORDER BY Date DESC, Time Desc;", username)
	if err != nil {
		// http status 500
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var thumbnail utilities.Thumbnail
	for rows.Next() {

		var date, time string
		errScan := rows.Scan(&thumbnail.Username, &thumbnail.PhotoId, &date, &time)
		if errScan != nil {
			return nil, fmt.Errorf("error while scanning Photo: %w", errScan)
		}
		thumbnail.DateTime = fmt.Sprintf("%sT%s", date, time)
		rowsLike := db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE Photo = ?;", thumbnail.PhotoId).Scan(&thumbnail.LikesNumber)
		if errors.Is(rowsLike, sql.ErrNoRows) {
			// http status 500
			return nil, fmt.Errorf("error execution query: %w", rowsLike)
		}
		rowsComment := db.c.QueryRow("SELECT COUNT(*) FROM Comment WHERE Photo = ?;", thumbnail.PhotoId).Scan(&thumbnail.CommentsNumber)
		if errors.Is(rowsComment, sql.ErrNoRows) {
			// http status 500
			return nil, fmt.Errorf("error execution query: %w", rowsComment)
		}
		thumbnail.PhotoURL = utilities.CreatePhotoURL(thumbnail.PhotoId)
		thumbnails = append(thumbnails, thumbnail)
	}

	errScan := rows.Err()
	if errScan != nil {
		return nil, fmt.Errorf("Error while scanning for GetThumbnails operation: %w", errScan)
	}

	// http status 200
	return thumbnails, nil

}
