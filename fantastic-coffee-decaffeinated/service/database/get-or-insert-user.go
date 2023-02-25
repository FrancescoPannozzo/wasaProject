package database

func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	result, err := DBcon.GetIdByName(name)

	if err != nil {
		result, err = DBcon.InsertUser(name)
	}

	return result, err
}
