package database

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Modify the user in the DB
func (db *appdbimpl) ModifyUsername(userid string, newName string) error {

	logrus.Info("Modifing the username in the db..")

	sqlStmt := fmt.Sprintf("UPDATE User SET Nickname = '%s' WHERE Id_user = '%s'", newName, userid)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("error execution query: %w", err)
	}

	logrus.Info("Modification done!")
	return nil
}
