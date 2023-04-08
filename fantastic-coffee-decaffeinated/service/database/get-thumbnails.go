package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"
)

func (db *appdbimpl) GetThumbnails(username string) ([]utilities.Thumbnail, error, int) {
	//var idphoto string

	var thumbnails []utilities.Thumbnail

	rows, err := db.c.Query("SELECT Id_photo, Date, Time FROM Photo WHERE User=? ORDER BY Date DESC, Time Desc;", username)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	var thumbnail utilities.Thumbnail
	for rows.Next() {

		var date, time string
		rows.Scan(&thumbnail.PhotoId, &date, &time)
		thumbnail.DateTime = fmt.Sprintf("%sT%s", date, time)
		rows := db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE Photo = ?;", thumbnail.PhotoId).Scan(&thumbnail.LikesNumber)
		if errors.Is(rows, sql.ErrNoRows) {
			return nil, fmt.Errorf("error execution query: %w", rows), http.StatusInternalServerError
		}
		rows = db.c.QueryRow("SELECT COUNT(*) FROM Comment WHERE Photo = ?;", thumbnail.PhotoId).Scan(&thumbnail.CommentsNumber)
		if errors.Is(rows, sql.ErrNoRows) {
			return nil, fmt.Errorf("error execution query: %w", rows), http.StatusInternalServerError
		}
		thumbnails = append(thumbnails, thumbnail)
		//calcola like number, comment number e selezionare datatime
		fmt.Println(thumbnail)
	}

	return thumbnails, nil, http.StatusOK

}