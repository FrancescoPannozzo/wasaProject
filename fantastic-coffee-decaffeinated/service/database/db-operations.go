package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func (db *appdbimpl) GetIdByName(name string) (string, error) {

	var (
		userID   string
		username string
	)

	//fmt.Println("now in CheckUser, name value is:", name)

	logrus.Infoln("now in GetIdByName(), name value is:", name)

	rows := db.c.QueryRow("SELECT Id_user, Nickname FROM User WHERE Nickname=?", name).Scan(&userID, &username)
	if errors.Is(rows, sql.ErrNoRows) {
		var errUser error
		userID, errUser = DBcon.InsertUser(name)
		return userID, errUser
	}

	fmt.Printf("User: %s already in the db\n", username)
	return userID, nil
}

func (db *appdbimpl) InsertUser(name string) (string, error) {
	//create user token
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(1000)
	s := strconv.Itoa(n)
	token := name + s

	fmt.Printf("token created =:%s", token)
	fmt.Println("Inserting the new user in the db..")

	sqlStmt := fmt.Sprintf("INSERT INTO User (Id_user, Nickname) VALUES('%s','%s');", token, name)

	fmt.Println(sqlStmt)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return "Query error", fmt.Errorf("error executin query: %w", err)
	}
	return token, nil
}

func (db *appdbimpl) GetOrInsertUser(name string) (string, error) {
	result, err := DBcon.GetIdByName(name)

	if err != nil {
		result, err = DBcon.InsertUser(name)
	}

	return result, err
}
