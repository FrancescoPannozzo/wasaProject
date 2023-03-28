package utilities

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
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
	var username Username
	err := json.NewDecoder(r.Body).Decode(&username)
	_ = r.Body.Close()
	if err != nil {
		logrus.Errorln("wrong JSON received")
		return "wrong JSON received", err
	}

	return username.Name, nil
}

// It sends a payload response json object with a http status code and a message related to it
func WriteResponse(httpStatus int, payload string, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(httpStatus)
	var response PayloadFeedback
	switch httpStatus {
	case 401, 400, 404, 500:
		response = &ErrorResponse{Error: payload}
	case 200, 201:
		response = &FeedbackResponse{Feedback: payload}
	}

	/*
		jsonResp, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	*/
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		logrus.Errorln("wrong JSON processed")
		json.NewEncoder(w).Encode(&response)
		return
	}
	//w.Write(jsonResp)
	//return
}

func CheckUsername(name string) (httpStatus int, feedback string) {
	if len(name) < 3 || len(name) > 13 {
		return http.StatusBadRequest, "Username not valid, size must be in range [3-13] characters"
	}
	return http.StatusOK, "Correct username type"
}

// Create an user id composed by characters + timestamp
func GenerateUserID(name string) string {
	//create user id
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var s string
	for i := 0; i < 4; i++ {
		n := r1.Intn(25) + 1
		s += string(toChar(n))
	}

	s += GenerateTimestamp()
	return s
}

func GenerateTimestamp() string {
	now := time.Now()
	return now.Format("20060102150405")
}

func toChar(i int) rune {
	return rune('a' - 1 + i)
}
