package database

import (
	"fmt"
	"net/http"
)

func (db *appdbimpl) UnbanUser(banner string, banned string) (string, error, int) {

	sqlStmt := fmt.Sprintf("DELETE FROM Ban WHERE Banner='%s' AND Banned='%s';", banner, banned)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return "error execution query", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "User unbanned, DB updated", err, http.StatusOK
}
