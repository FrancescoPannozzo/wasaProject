package database

// GetOrInsertUser check if a username exist in the DB, if exists return the user id and nil
// otherwise the user is stored in the DB.
// It returns an empty string and the error in unable to insert the user in the DB.
func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	userId, err := DBcon.GetIdByName(name)

	if err != nil {
		//user not present in the DB, interting attempt in the DB
		userId, err = DBcon.InsertUser(name)
		if err != nil {
			// err!= nil stands for 500 http status
			return "", err
		}
	}
	// Return a feedback message and nil if successfull.
	// Return a feedback message and an error excetution query otherwise.
	// err = nil stands for 201 http status created
	return userId, nil
}
