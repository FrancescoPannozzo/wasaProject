package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) CheckBan(loggedUser string, targetUser string) bool {
	var banned string
	rows := db.c.QueryRow("SELECT Banned FROM Ban WHERE Banner =? AND Banned =?;", targetUser, loggedUser).Scan(&banned)

	if errors.Is(rows, sql.ErrNoRows) {
		fmt.Printf("user %s is not banned\n", banned)
		return false
	}
	fmt.Printf("banned user: %s\n", banned)
	return true
}
