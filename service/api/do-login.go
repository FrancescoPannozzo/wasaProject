package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Do the loggin with the username provided in the requestBody
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Logging the user..")

	var username utilities.Username
	err := json.NewDecoder(r.Body).Decode(&username)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Warning("wrong JSON received")
		utilities.WriteResponse(http.StatusInternalServerError, "cannot read the request", w)
		return
	}

	errCheckUser := utilities.CheckUsername(username.Name)

	if err != nil {
		utilities.WriteResponse(http.StatusBadRequest, errCheckUser.Error(), w)
		return
	}

	testUserID, errUser := database.DBcon.GetOrInsertUser(username.Name)

	if errUser != nil {
		fmt.Printf("Cannot send the ID: %v\n", errUser)
		utilities.WriteResponse(http.StatusInternalServerError, fmt.Sprintf("Cannot send the ID: %v\n", errUser), w)
		return
	}

	type UserID struct {
		Identifier string `json:"identifier"`
	}

	userId := UserID{Identifier: testUserID}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	errEnc := json.NewEncoder(w).Encode(&userId)
	if errEnc != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errEnc.Error(), w)
		return
	}

	logrus.Infoln("User logged!")
	return
}
