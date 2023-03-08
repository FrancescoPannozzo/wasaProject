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

// A rappresentation of the username
type Username struct {
	Name string `json:"name"`
}

// A rappresentation of a 4XX message error in string format
type ErrorResponse struct {
	Error string `json:"error"`
}

type FeedbackResponse struct {
	Feedback string `json:"feedback"`
}

// Verify the user id from a request, return the http status number and the message related to it
func VerifyUseridController(w http.ResponseWriter, r *http.Request) (int, string) {
	prefix := "Baerer "
	authHeader := r.Header.Get(("Authorization"))
	log.Println(authHeader)

	reqUserid := strings.TrimPrefix(authHeader, prefix)
	log.Println(reqUserid)

	username := r.URL.Query().Get("username")

	userid, errUserid := database.DBcon.GetIdByName(username)

	if errUserid == nil && reqUserid == userid {
		return 200, "Successfull Authorization, Access allowed"
	}

	if errUserid != nil {
		fmt.Println(errUserid)
		return 400, "Error while retriving the username identifier from the DB"
	}

	if authHeader == "" || reqUserid == authHeader || reqUserid != userid {
		return 400, "User ID not valid"
	}

	return 400, "Error in authentication"
}

// Get the username from a request, return the username and nil.
// If an error has occurred return an empty string and the error
func GetNameFromReq(r *http.Request) (string, error) {
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

// It sent a payload response with a http status code and a message related to it
func WriteResponse(httpStatus int, payload string, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-type", "application/json")
	response := FeedbackResponse{Feedback: payload}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(jsonResp)
	return
}
