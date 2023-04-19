package database

import (
	"fmt"
)

// Modify the user in the DB
func (db *appdbimpl) ModifyUsername(userid string, newName string) error {
	sqlStmt := fmt.Sprintf("UPDATE User SET Nickname = '%s' WHERE Id_user = '%s'", newName, userid)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("error execution query: %w", err)
	}

	return nil
}
