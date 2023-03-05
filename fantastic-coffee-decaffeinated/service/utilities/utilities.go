package utilities

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Verify the token from a request
func VerifyTokenController(w http.ResponseWriter, r *http.Request) bool {
	prefix := "Baerer "
	authHeader := r.Header.Get(("Authorization"))
	log.Println(authHeader)

	reqToken := strings.TrimPrefix(authHeader, prefix)
	log.Println(reqToken)

	username, errName := GetUserFromReq(r)

	if errName != nil {
		errToken := fmt.Errorf("error while reading the body request: %w", errName)
		fmt.Println(errToken)
		return false
	}

	token, errToken := database.DBcon.GetIdByName(username)

	if errToken != nil {
		fmt.Println(errToken)
		return false
	}

	if authHeader == "" || reqToken == authHeader || reqToken != token {
		//return error4XX and payload message
		return false
	}

	if reqToken == token {
		//return response 2XX and payload message
		return true
	}

	return true
}

type Username struct {
	Name string `json:"name"`
}

// Get the username from a request
func GetUserFromReq(r *http.Request) (string, error) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errBody := fmt.Errorf("error while reading the body request: %v", err)
		fmt.Println(errBody)
		return "", errBody
	}
	var username Username
	errConv := json.Unmarshal(reqBody, &username)

	if errConv != nil {
		fmt.Printf("error with unmarshal.. err: %v", errConv)
		return "", errConv
	}

	return username.Name, nil
}
