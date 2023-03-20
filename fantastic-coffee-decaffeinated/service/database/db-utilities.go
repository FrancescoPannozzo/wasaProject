package database

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Verify the user id from a request with a Baerer Authorization Header, return the http status number and the message related to it
func VerifyUseridController(w http.ResponseWriter, r *http.Request) (int, string) {
	prefix := "Baerer "
	authHeader := r.Header.Get(("Authorization"))
	log.Println(authHeader)

	reqUserid := strings.TrimPrefix(authHeader, prefix)
	log.Println(reqUserid)

	username := r.URL.Query().Get("username")

	// Searching the username in the database
	userid, errUserid := DBcon.GetIdByName(username)

	if errUserid != nil {
		fmt.Println(errUserid)
		return http.StatusBadRequest, "Error while retriving the username identifier from the DB"
	}

	if errUserid == nil && reqUserid == userid {
		return http.StatusOK, "Successfull Authorization, Access allowed"
	}

	if authHeader == "" || reqUserid == authHeader || reqUserid != userid {
		return http.StatusUnauthorized, "User ID not valid"
	}

	return http.StatusBadRequest, "Error in authentication"
}
