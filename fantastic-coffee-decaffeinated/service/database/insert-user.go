package database

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Insert the user in the DB,
// return the userID, error = nil and the the http.StatusCreated
// If there is an internal error it return empty userID, error != nil and http.StatusInternalServerError
func (db *appdbimpl) InsertUser(name string) (string, error, int) {
	//create user id
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(1000)
	s := strconv.Itoa(n)
	userID := name + s

	logrus.Info("User id created =%s, inserting the new user in the db..\n", userID)

	sqlStmt := fmt.Sprintf("INSERT INTO User (Id_user, Nickname) VALUES('%s','%s');", userID, name)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		// 500 Internal server error
		return "", fmt.Errorf("error execution query: %w", err), http.StatusInternalServerError
	}

	// 201 Created
	return userID, nil, http.StatusCreated
}
