package database

// GetOrInsertUser check if a username exist in the DB, if exists return the user id,
// otherwise the user is stored in the DB
func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	userId, err := DBcon.GetIdByName(name)

	if err != nil {
		//user not present in the DB, interting attempt in the DB
		//var httpResponse int
		userId, err = DBcon.InsertUser(name)
		if err != nil {
			// return 500 http status
			return "", err
		}
		// return 201 http status created
		//return userId, err, httpResponse
	}
	// Insert the comment on the photoID provided in the DB.
	// Return a feedback message and nil if successfull.
	// Return a feedback message and an error excetution query otherwise.
	// return 201 http status created
	return userId, nil
}
