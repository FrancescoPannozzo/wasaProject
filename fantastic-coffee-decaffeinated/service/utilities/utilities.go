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

type Username struct {
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type FeedbackResponse struct {
	Feedback string `json:"feedback"`
}

// Verify the user id from a request
func VerifyUseridController(w http.ResponseWriter, r *http.Request) (int, string) {
	prefix := "Baerer "
	authHeader := r.Header.Get(("Authorization"))
	log.Println(authHeader)

	reqUserid := strings.TrimPrefix(authHeader, prefix)
	log.Println(reqUserid)

	username := r.URL.Query().Get("username")

	fmt.Println("In VerifyUseridController(), username parameter is:", username)

	userid, errUserid := database.DBcon.GetIdByName(username)

	if errUserid == nil && reqUserid == userid {
		return 200, "Successfull Authorization, Access allowed"
	}

	if errUserid != nil {
		fmt.Println(errUserid)
		return 400, "Error while retriving the username identifier from the DB"
	}

	if authHeader == "" || reqUserid == authHeader || reqUserid != userid {
		//return error4XX and payload message
		return 400, "User ID not valid"
	}

	if reqUserid == userid {
		//return response 2XX and payload message
		return 200, "Successfull Authorization, Access allowed"
	}

	return 400, "Error, cannot detect the error origin"
}

// Get the username from a request
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
