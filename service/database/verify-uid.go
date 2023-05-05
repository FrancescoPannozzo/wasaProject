package database

import (
	"errors"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Verify the user id from a request with a Baerer Authorization Header, return the http status number and the message related to it
func VerifyUserId(r *http.Request, ps httprouter.Params) error {
	authHeader := r.Header.Get("Authorization")
	baererUserID := utilities.GetBearerID(r)

	if authHeader == "" || baererUserID == authHeader {
		return errors.New("Bearer Token not valid")
	}

	// Searching the username in the database
	_, errNickname := DBcon.GetNameByID(baererUserID)

	if errNickname != nil {
		return errors.New("Unauthorized user")
	}

	if errNickname == nil {
		return nil
	}

	return errors.New("Unauthorized user")
}
