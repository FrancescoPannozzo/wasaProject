package database

// GetOrInsertUser check if a username exist in the DB, if exists return the user id
// else the user is inserted in the DB
func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	result, err := DBcon.GetIdByName(name)

	if err != nil {
		result, err = DBcon.InsertUser(name)
	}

	return result, err
}
