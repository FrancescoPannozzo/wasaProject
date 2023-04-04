package database

import (
	"fmt"
	"net/http"
)

func (db *appdbimpl) BanUser(banner string, banned string) (string, error, int) {

	sqlStmt := fmt.Sprintf("INSERT INTO Ban (Banner, Banned) VALUES('%s','%s');", banner, banned)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return "error execution query", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	return "Banned user inserted in the DB", err, http.StatusCreated
}
