package database

// GetOrInsertUser check if a username exist in the DB, if exists return the user id
// else the user is inserted in the DB
func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	// var httpStatus int
	userId, err := DBcon.GetIdByName(name)

	if err != nil {
		userId, err = DBcon.InsertUser(name)
		if err != nil {
			// return "", err, 404
			return "", err
		}
		//RITORNA ANCHE HTTP STATUS CODE!!
		// httpStatus = 201
	}
	//httpStatus = 200
	// return "", err, 200
	return userId, err
}
