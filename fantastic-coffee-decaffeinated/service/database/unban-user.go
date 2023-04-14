package database

import (
	"fmt"
)

func (db *appdbimpl) UnbanUser(banner string, banned string) (string, error) {

	sqlStmt := fmt.Sprintf("DELETE FROM Ban WHERE Banner='%s' AND Banned='%s';", banner, banned)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		//500
		return "error execution query", fmt.Errorf("error execution query: %w", err)
	}

	//200
	return "User unbanned, DB updated", err
}
