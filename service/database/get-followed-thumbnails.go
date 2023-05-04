package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

func (db *appdbimpl) GetFollowedThumbnails(loggedUser string) ([]utilities.Thumbnail, error) {
	var thumbnails []utilities.Thumbnail

	rows, err := db.c.Query("SELECT Id_photo, User, Date, Time FROM Photo JOIN Follow ON Follow.Followed = Photo.User WHERE Follow.Follower =? ORDER BY Date DESC, Time Desc;", loggedUser)
	if err != nil {
		// http status 500
		return nil, fmt.Errorf("error execution query: %w", err)
	}

	var thumbnail utilities.Thumbnail
	for rows.Next() {
		var idphoto, user, date, time string
		errScan := rows.Scan(&idphoto, &user, &date, &time)
		if errScan != nil {
			return nil, fmt.Errorf("error while scanning Follow: %w", errScan)
		}
		if db.CheckBan(loggedUser, user) {
			continue
		}
		thumbnail.Username = user
		thumbnail.PhotoId = idphoto
		thumbnail.PhotoURL = utilities.CreatePhotoURL(idphoto)
		thumbnail.DateTime = fmt.Sprintf("%sT%s", date, time)
		rowsLike := db.c.QueryRow("SELECT COUNT(*) FROM Like WHERE Photo = ?;", idphoto).Scan(&thumbnail.LikesNumber)
		if errors.Is(rowsLike, sql.ErrNoRows) {
			return nil, fmt.Errorf("error execution query: %w", rowsLike)
		}
		rowsComment := db.c.QueryRow("SELECT COUNT(*) FROM Comment WHERE Photo = ?;", idphoto).Scan(&thumbnail.CommentsNumber)
		if errors.Is(rowsComment, sql.ErrNoRows) {
			return nil, fmt.Errorf("error execution query: %w", rowsComment)
		}
		thumbnails = append(thumbnails, thumbnail)
	}

	// http status 200
	return thumbnails, nil
}
