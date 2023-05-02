package database

// Verify the username to change ownership
func (db *appdbimpl) CheckOwnership(userId string, username string) bool {
	usernameCheck, err := db.GetNameByID(userId)

	if err != nil {
		return false
	}

	return usernameCheck == username
}
