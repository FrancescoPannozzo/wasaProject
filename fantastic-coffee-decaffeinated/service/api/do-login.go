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

	err = utilities.CheckUsername(username.Name)

	if err != nil {
		utilities.WriteResponse(http.StatusBadRequest, err.Error(), w)
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
	w.Header().Set("content-type", "application/json") //si setta quello che mandi
	err = json.NewEncoder(w).Encode(&userId)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	logrus.Infoln("..user logged!")

	return
}
