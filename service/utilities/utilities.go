package utilities

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// ----- ENTITIES ------

const Unauthorized = "Unauthorized user"
const ErrorExecutionQuery = "error execution query in DB"

// A rappresentation of a thumbnail image with informations
type Thumbnail struct {
	Username       string `json:"username"`
	PhotoId        string `json:"photoid"`
	PhotoURL       string `json:"photourl"`
	DateTime       string `json:"datetime"`
	LikesNumber    int    `json:"nlikes"`
	CommentsNumber int    `json:"ncomments"`
}

// A rappresentation of the username
type Username struct {
	Name string `json:"name"`
}

// A rappresentation of a comment
type Comment struct {
	CommentId string `json:"commentid"`
	Name      string `json:"name"`
	Content   string `json:"comment"`
}

// a rappresentation of a thumbnail image with informations
type Post struct {
	Username       string    `json:"username"`
	PhotoURL       string    `json:"photourl"`
	DateTime       string    `json:"datetime"`
	LikesNumber    int       `json:"nlikes"`
	Comments       []Comment `json:"comments"`
	LikeThis       bool      `json:"likethis"`
	LoggedUsername string    `json:"loggedusername"`
}

// a rappresentation of a user profile
type Profile struct {
	VisitedUsername   string      `json:"visitedUsername"`
	LoggedUsername    string      `json:"loggedUsername"`
	PhotoNumber       int         `json:"nphoto"`
	Followers         []string    `json:"followers"`
	Followed          []string    `json:"followed"`
	Thumbnail         []Thumbnail `json:"thumbnails"`
	LoggedUserBanned  bool        `json:"loggedUserBanned"`
	VisitedUserBanned bool        `json:"visitedUserBanned"`
}

// a rappresentation of a my-stream page with informations
type MyStream struct {
	LoggedUsername string      `json:"loggedUsername"`
	Thumbnails     []Thumbnail `json:"thumbnails"`
}

// A feedback message rappresentation in string format
type FeedbackResponse struct {
	Feedback string `json:"feedback"`
}

// Custom Error type
type DbBadRequestError struct{}

// ----- FUNCTIONS -----

func (e *DbBadRequestError) Error() string {
	return "Bad request for the provided DB operation"
}

// Get the username from a request, return the username and nil if successful.
// If an error has occurred return a feedback string and the error
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

	response := &FeedbackResponse{Feedback: payload}

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		logrus.Errorln("wrong JSON processed")
		message := "{\"error\":\"wrong JSON processed\"}"
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte(message))
		if errWrite != nil {
			logrus.Warn("Cannot write in the ResponseWriter")
			return
		}
		return
	}
}

// Checks if the provided username length is valid for the application.
// Returns nil if successfull, an error otherwise
func CheckUsername(name string) error {
	if len(name) < 3 || len(name) > 13 {
		return errors.New("Username not valid, size must be in range [3-13] characters")
	}
	return nil
}

// Create a user id composed by random characters + timestamp
func GenerateUserID() string {
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

// Photo URL maker
func CreatePhotoURL(idPhoto string) string {
	baseURL := "http://0.0.0.0:3000/photos/"
	return filepath.Join(baseURL, idPhoto)
}

// Check if the photo ID format length is valid
func IsPhotoIdValid(idphoto string) bool {
	data := strings.Split(idphoto, ".")
	const idphotoLenghts = 18
	return len(data[0]) == idphotoLenghts
}
