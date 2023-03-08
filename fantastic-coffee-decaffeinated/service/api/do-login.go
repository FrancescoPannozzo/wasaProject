package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Logging the user..")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errBody := fmt.Errorf("error while reading the body request: %v", err)
		fmt.Println(errBody)
		return
	}

	//fmt.Println("reqBody content is:", bytes.NewBuffer(reqBody).String())

	//isValid := json.Valid(reqBody)
	//fmt.Println("Is reqBody content a valid json format?: ", isValid)

	type Username struct {
		Name string `json:"name"`
	}

	var username Username
	errConv := json.Unmarshal(reqBody, &username)

	if errConv != nil {
		fmt.Printf("error with unmarshal.. err: %v", errConv)
		return
	}

	//fmt.Println("username to store in db is:", username.Name)

	testToken, errUser := database.DBcon.GetOrInsertUser(username.Name)

	if errUser != nil {
		fmt.Printf("Cannot sent the ID, error: %v\n", errUser)
		return
	}

	//fmt.Printf("Token for user %s is %s\n", username.Name, testToken)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json") //si setta quello che mandi

	type UserID struct {
		Identifier string `json:"identifier"`
	}

	userId := UserID{testToken}

	jsonResp, errJson := json.Marshal(&userId)
	if err != nil {
		logrus.Infof("Error with Marshal: %v\n", errJson)
		return
	}

	logrus.Infoln("..user logged!")
	w.Write(jsonResp)

	return
}
