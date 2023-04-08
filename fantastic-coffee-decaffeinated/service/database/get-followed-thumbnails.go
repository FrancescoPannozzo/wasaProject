package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"
)

func (db *appdbimpl) GetFollowedThumbnails(loggedUser string) ([]utilities.Thumbnail, error, int) {
	//var idphoto string

	var thumbnails []utilities.Thumbnail

	rows, err := db.c.Query("SELECT Id_photo, User, Date, Time FROM Photo JOIN Follow ON Follow.Followed = Photo.User WHERE Follow.Follower =? ORDER BY Date DESC, Time Desc;", loggedUser)
	if err != nil {
		return nil, fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	var thumbnail utilities.Thumbnail
	for rows.Next() {

		var idphoto, user, date, time string
		rows.Scan(&idphoto, &user, &date, &time)
		if db.CheckBan(loggedUser, user) {
			continue
		}
		thumbnail.PhotoId = idphoto
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
