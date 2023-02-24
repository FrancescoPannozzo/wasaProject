package api

import (
	"bytes"
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json") //si setta quello che mandi

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errBody := fmt.Errorf("error while reading the body request: %v", err)
		fmt.Println(errBody)
		return
	}
	fmt.Println("reqBody content is:", bytes.NewBuffer(reqBody).String())
	_, _ = w.Write(reqBody)

	isValid := json.Valid(reqBody)
	fmt.Println("Is reqBody content a valid json format?: ", isValid)

	type Username struct {
		Name string `json:"name"`
	}

	var username Username
	errConv := json.Unmarshal(reqBody, &username)

	if errConv != nil {
		fmt.Printf("error with unmarshal.. err: %v", errConv)
		return
	}

	fmt.Println("username to store in db is:", username.Name)

	testToken, _ := database.DBcon.GetOrInsertUser(username.Name)

	fmt.Printf("Token for user %s is %s", username.Name, testToken)

}
