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
	statusNumber, payloadMessage := utilities.VerifyUseridController(w, r)

	if statusNumber == http.StatusBadRequest {
		utilities.WriteResponse(http.StatusBadRequest, payloadMessage, w)
		return
	}

	oldUsername := r.URL.Query().Get("username")

	newUsername, errName := utilities.GetNameFromReq(r)

	if errName != nil {
		logrus.Infof("Error in setMyUsername() while getting the username from the client request %v", errName)
		return
	}

	userid, errDb, httpResponse := database.DBcon.GetIdByName(oldUsername)

	if errDb != nil {
		logrus.Infof("Error in setMyUsername() while getting the user id from the client request %v", errName)
		return
	}

	err := database.DBcon.ModifyUsername(userid, newUsername)

	if err != nil {
		fmt.Println(err)
		utilities.WriteResponse(httpResponse, err.Error(), w)
		return
	}

	utilities.WriteResponse(httpResponse, "Username successfully updated", w)
	return

	//verifica l'auth, verificandol'auth ottengo token, estrapola nuovo username dalla
	// request, cerca record con il token e modifica lo username ad esso associato
}
