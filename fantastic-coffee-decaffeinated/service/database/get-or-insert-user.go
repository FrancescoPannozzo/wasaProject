package database

import "net/http"

// GetOrInsertUser check if a username exist in the DB, if exists return the user id,
// else the user is stored in the DB
func (db *appdbimpl) GetOrInsertUser(name string) (string, error, int) {
	userId, err := DBcon.GetIdByName(name)

	if err != nil {
		//user not present in the DB, interting attempt in the DB
		var httpResponse int
		userId, err, httpResponse = DBcon.InsertUser(name)
		if err != nil {
			// return 500 http status
			return "", err, httpResponse
		}
		// return 201 http status created
		//return userId, err, httpResponse
	}

	// return 201 http status created
	return userId, err, http.StatusCreated
}
