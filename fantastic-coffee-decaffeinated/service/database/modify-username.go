package database

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Modify the user in the DB
func (db *appdbimpl) ModifyUsername(userid string, newName string) error {

	logrus.Info("Updating the username in the db..")

	newUserID := utilities.GenerateUserID(newName)

	sqlStmt := fmt.Sprintf("UPDATE User SET Nickname = '%s', Id_user = '%s' WHERE Id_user = '%s'", newName, newUserID, userid)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return fmt.Errorf("error execution query: %w", err)
	}

	logrus.Info("Update done!")
	return nil
}
