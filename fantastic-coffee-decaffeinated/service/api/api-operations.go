package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"bytes"
	"fmt"
	"io"
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
)


func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errBody := fmt.Errorf("error while reading the body request: %v", err)
		fmt.Println(errBody)
		return
	}
	fmt.Println("reqBody content is:/n", bytes.NewBuffer(reqBody).String())
	_, _ = w.Write(reqBody)

	isValid := json.Valid(reqBody)
	fmt.Println("/nIs reqBody content a valid json format?: ", isValid)

	type Username struct{
		Name string `json:"name"`
	}

	var username Username
	errConv := json.Unmarshal(reqBody, &username)
	fmt.Println("errConv is:", errConv)
	if errConv != nil {
		fmt.Println("error with unmarshal.. err: %v", errConv)
		return
	}

	fmt.Println("username to store in db is:", username.Name)

	testUsername, _ := database.DBcon.CheckUser(username.Name)

	fmt.Println("Username from CheckUser:", testUsername)

}