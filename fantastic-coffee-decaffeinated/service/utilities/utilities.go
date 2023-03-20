package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

// Create an user id of type 'username + a random number'
func GenerateUserID(name string) string {
	//create user id
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(1000)
	s := strconv.Itoa(n)
	return name + s
}
