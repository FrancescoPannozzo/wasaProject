package database

import (
	"fmt"
)

// Ban the provided user. Returns a feedback string and nil if successful,
// a feedback string and error != nil otherwise.
func (db *appdbimpl) BanUser(banner string, banned string) (string, error) {

	sqlStmt := fmt.Sprintf("INSERT INTO Ban (Banner, Banned) VALUES('%s','%s');", banner, banned)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		//500
		return "error execution query", fmt.Errorf("error execution query: %w", err)
	}
	//201
	return "Banned user inserted in the DB", nil
}
