package utilities

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// A rappresentation of the username
type Username struct {
	Name string `json:"name"`
}

type PayloadFeedback interface {
	PrintFeedback(s string) string
}

// A rappresentation of a 4XX message error in string format
type ErrorResponse struct {
	Error string `json:"error"`
}

// A rappresentation of a 2XX message error in string format
type FeedbackResponse struct {
	Feedback string `json:"feedback"`
}

func (er *ErrorResponse) PrintFeedback(s string) string {
	response := fmt.Sprint("Error response:", s)
	return response
}

func (fr *FeedbackResponse) PrintFeedback(s string) string {
	response := fmt.Sprint("Feedback response:", s)
	return response
}

// Verify the user id from a request with a Baerer Authorization Header, return the http status number and the message related to it
func VerifyUseridController(w http.ResponseWriter, r *http.Request) (int, string) {
	prefix := "Baerer "
	authHeader := r.Header.Get(("Authorization"))
	//log.Println(authHeader)

	reqUserid := strings.TrimPrefix(authHeader, prefix)
	//log.Println(reqUserid)

	username := r.URL.Query().Get("username")

	// Searching the username in the database
	userid, errUserid, _ := database.DBcon.GetIdByName(username)

	if errUserid == nil && reqUserid == userid {
		return http.StatusOK, "Successfull Authorization, Access allowed"
	}

	if errUserid != nil {
		fmt.Println(errUserid)
		return http.StatusBadRequest, "Error while retriving the username identifier from the DB"
	}

	if authHeader == "" || reqUserid == authHeader || reqUserid != userid {
		return http.StatusUnauthorized, "User ID not valid"
	}

	return http.StatusBadRequest, "Error in authentication"
}

// Get the username from a request, return the username and nil if successful.
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

// It send a payload response json object with a http status code and a message related to it
func WriteResponse(httpStatus int, payload string, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-type", "application/json")
	var response PayloadFeedback
	switch httpStatus {
	case 401, 400, 404, 500:
		response = &ErrorResponse{Error: payload}
	case 200, 201:
		response = &FeedbackResponse{Feedback: payload}
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(jsonResp)
	return
}

func CheckUsername(name string) (httpStatus int, feedback string) {
	if len(name) < 3 || len(name) > 13 {
		return http.StatusBadRequest, "Username not valid, size must be in range [3-13] characters"
	}
	return http.StatusOK, "Correct username type"
}
