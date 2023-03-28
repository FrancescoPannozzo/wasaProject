package database

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Verify the user id from a request with a Baerer Authorization Header, return the http status number and the message related to it
func VerifyUseridController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (int, string) {
	authHeader := r.Header.Get(("Authorization"))
	baererUserID := GetBaererID(r)

	// Searching the username in the database
	_, errNickname := DBcon.GetNameByID(baererUserID)

	if errNickname != nil || authHeader == "" || baererUserID == authHeader {
		fmt.Println(errNickname)
		return http.StatusUnauthorized, "Baerer Token not valid"
	}

	if errNickname == nil {
		return http.StatusOK, "Successfull Authorization, Access allowed"
	}

	return http.StatusBadRequest, "Error in authentication"
}

func GetBaererID(r *http.Request) string {
	prefix := "Baerer "
	authHeader := r.Header.Get(("Authorization"))
	log.Println(authHeader)
	return strings.TrimPrefix(authHeader, prefix)
}
