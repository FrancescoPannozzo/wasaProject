package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Update an existing username
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	statusNumber, payloadMessage := database.VerifyUseridController(w, r)

	if statusNumber == http.StatusBadRequest || statusNumber == http.StatusUnauthorized {
		logrus.Infof("Error with the authentication, httpStatus is '%v', %s", statusNumber, payloadMessage)
		utilities.WriteResponse(statusNumber, payloadMessage, w)
		return
	}

	oldUsername := r.URL.Query().Get("username")

	newUsername, errName := utilities.GetNameFromReq(r)

	if errName != nil {
		logrus.Infof("Error in setMyUsername() while getting the username from the client request %v", errName)
		return
	}

	statusNumber, payloadMessage = utilities.CheckUsername(newUsername)
	if statusNumber == http.StatusBadRequest {
		utilities.WriteResponse(statusNumber, payloadMessage, w)
		return
	}

	userid, errDb := database.DBcon.GetIdByName(newUsername)
	if errDb == nil {
		message := fmt.Sprintf("WARNING, the username %s is already taken, please choose another one", newUsername)
		logrus.Println(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	userid, errDb = database.DBcon.GetIdByName(oldUsername)

	if errDb != nil {
		// POTENZIALE ERRORE 500 INTERNAL SERVER ERROR
		logrus.Infof("Error in setMyUsername() while getting the user id from the client request %v", errName)
		utilities.WriteResponse(http.StatusInternalServerError, "Error while getting the user id from the client request", w)
		return
	}

	err := database.DBcon.ModifyUsername(userid, newUsername)

	if err != nil {
		fmt.Println(err)
		//OCCHIO QUI A RITORNARE GLI HTTP STATUS NUMBERS - AGGIORNARE API
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}
	//OCCHIO QUI A RITORNARE GLI HTTP STATUS NUMBERS
	utilities.WriteResponse(http.StatusCreated, "Username successfully updated", w)
	return

	//verifica l'auth, verificandol'auth ottengo token, estrapola nuovo username dalla
	// request, cerca record con il token e modifica lo username ad esso associato
}
