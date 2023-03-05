package database

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Insert the user in the DB
func (db *appdbimpl) InsertUser(name string) (string, error) {
	//create user id
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(1000)
	s := strconv.Itoa(n)
	token := name + s

	logrus.Info("Inserting the new user in the db..")
	logrus.Infof("User id created =%s\n", token)

	sqlStmt := fmt.Sprintf("INSERT INTO User (Id_user, Nickname) VALUES('%s','%s');", token, name)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return "query-error", fmt.Errorf("error execution query: %w", err)
	}

	return token, nil
}
