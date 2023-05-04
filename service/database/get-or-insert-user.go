package database

// GetOrInsertUser check if a username exist in the DB, if exists return the user id and nil
// otherwise the user is stored in the DB returning the userID.
// It returns an empty string and the error in unable to insert the user in the DB.
func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	userId, err := DBcon.GetIdByName(name)

	if err != nil {
		// user not present in the DB, interting attempt in the DB
		userId, errInsUser := DBcon.InsertUser(name)
		if errInsUser != nil {
			// err!= nil stands for 500 http status
			return "", err
		}
		return userId, errInsUser
	}

	return userId, nil
}
