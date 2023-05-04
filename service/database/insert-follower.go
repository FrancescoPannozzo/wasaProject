package database

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
)

// Insert the user in the DB,
func (db *appdbimpl) InsertFollower(follower string, followed string) (string, error) {
	var user string
	rows := db.c.QueryRow("SELECT Follower FROM Follow WHERE Follower=? AND Followed=?", follower, followed).Scan(&user)
	if !errors.Is(rows, sql.ErrNoRows) {
		return "Warning, the user already follow the target user", &utilities.DbBadRequestError{}
	}

	sqlStmt := fmt.Sprintf("INSERT INTO Follow (Follower, Followed) VALUES('%s','%s');", follower, followed)
	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return utilities.ErrorExecutionQuery, fmt.Errorf("error execution query: %w", err)
	}

	// http status 201
	return "follow added, ok", nil

}
