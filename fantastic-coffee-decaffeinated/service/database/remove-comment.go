package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

// Delete a follow
func (db *appdbimpl) RemoveComment(username string, idphoto string, idcomment string) (string, error) {

	var content string

	rows := db.c.QueryRow("SELECT Content FROM Comment WHERE User = '?' AND Id_comment = ? AND Photo = '?';", username, idcomment, idphoto).Scan(&content)

	if errors.Is(rows, sql.ErrNoRows) {
		logrus.Printf("no results")
		//400
		return "no results", rows
	}

	log.Printf("%v - %v - %v", username, idphoto, idcomment)
	sqlStmt := fmt.Sprintf("DELETE FROM Comment WHERE User = '%s' AND Id_comment = %s AND Photo = '%s';", username, idcomment, idphoto)
	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "error execution query in DB", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "comment removed, ok", nil

}
