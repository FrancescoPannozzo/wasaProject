package database

// Insert the user in the DB,
func (db *appdbimpl) LikePhoto(username string, idphoto string) (string, error) {
	_, err := db.c.Exec("INSERT INTO Like (User, Photo) VALUES(?, ?);", username, idphoto)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", err
	}
	//201
	return "like added, ok", nil
}
