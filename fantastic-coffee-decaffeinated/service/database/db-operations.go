package database

import (
	"fmt"
	"database/sql"
	"errors"
)

// CheckUser checks if one user is in the db, if present it return the user authorization identifier
// for header purpose, if not present it create a new record with username and user identifier
func (db *appdbimpl) CheckUser(name string) (string, error) {
	
	var (
		userID string
		username string
	)

	fmt.Println("now in CheckUser, name value is:", name)

	rows := db.c.QueryRow("SELECT Nickname FROM User WHERE Nickname=?",name).Scan(&userID, &username)
	
	if errors.Is(rows, sql.ErrNoRows) {
		fmt.Println("Inserting the new user in the db..")
		return "INSERTED", nil
		/*
		errName := fmt.Errorf("error while reading the body request: %v", errName)
		fmt.println("error, no rows found:%v", errName)
		return nil, errName
		*/
	}

	/*
	err := rows.Scan(&userID, &username)
	if  err!= nil {
		errF := fmt.Errorf("errore nella lettura del record")
		fmt.Println(errF)
	}
	*/
	return "ALREADY IN DB", nil
}
