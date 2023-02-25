package database

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func (db *appdbimpl) InsertUser(name string) (string, error) {
	//create user token
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(1000)
	s := strconv.Itoa(n)
	token := name + s

	fmt.Printf("token created =:%s\n", token)
	fmt.Println("Inserting the new user in the db..")

	sqlStmt := fmt.Sprintf("INSERT INTO User (Id_user, Nickname) VALUES('%s','%s');", token, name)

	fmt.Println(sqlStmt)

	_, err := db.c.Exec(sqlStmt)

	if err != nil {
		return "query-error", fmt.Errorf("error execution query: %w", err)
	}
	return token, nil
}
