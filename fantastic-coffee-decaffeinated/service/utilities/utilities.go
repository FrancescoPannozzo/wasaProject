package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// a rappresentation of a thubnail image with informations
type Thumbnail struct {
	PhotoId        string `json:"idphoto"`
	DateTime       string `json:"datetime"`
	LikesNumber    int    `json:"nlikes"`
	CommentsNumber int    `json:"ncomments"`
}

// A rappresentation of the username
type Username struct {
	Name string `json:"name"`
}

// An Interfece to be able to set the right type of payload message, error or feedback
type PayloadFeedback interface {
	PrintFeedback(s string) string
}

// @todo: model, entities -> comment.go
// A rappresentation of a comment
type Comment struct {
	Name    string `json:"name"`
	Content string `json:"comment"`
}

// A rappresentation of a 4XX/500 message error in string format
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
	// @todo: mettere i range invece dei numeri
	switch httpStatus {
	case 401, 400, 404, 500:
		response = &ErrorResponse{Error: payload}
	// @todo: inutile
	case 200, 201:
		response = &FeedbackResponse{Feedback: payload}
	}

	// @todo: mandare solo stringa
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		logrus.Errorln("wrong JSON processed")
		json.NewEncoder(w).Encode("{\"error\":\"wrong JSON processed\"}")
		return
	}
}

// Checks if the username provided is valid for the application.
// Returns nil if successfull, an error otherwise
func CheckUsername(name string) error {
	if len(name) < 3 || len(name) > 13 {
		return errors.New("Username not valid, size must be in range [3-13] characters")
	}
	return nil
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

// get the bearer id from the requestBody
func GetBearerID(r *http.Request) string {
	prefix := "Bearer "
	authHeader := r.Header.Get(("Authorization"))
	return strings.TrimPrefix(authHeader, prefix)
}
