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
	logrus.Infoln("Updating the username in the db..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		logrus.Warn(errId.Error())
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	// Get the username to change from the RequestBody
	newUsername, errName := utilities.GetNameFromReq(r)
	if errName != nil {
		logrus.Warn("Error while getting the username from the client request")
		utilities.WriteResponse(http.StatusBadRequest, "Error: requestBody not valid", w)
		return
	}

	errCheckUser := utilities.CheckUsername(newUsername)
	if errCheckUser != nil {
		logrus.Warn(errCheckUser.Error())
		utilities.WriteResponse(http.StatusBadRequest, errCheckUser.Error(), w)
		return
	}

	// Check if who wants to change the username is the real profile owner
	loggedUserId := utilities.GetBearerID(r)
	oldUsername := ps.ByName("username")

	if !database.DBcon.CheckOwnership(loggedUserId, oldUsername) {
		feedback := "attempt to change someone else's username detected"
		logrus.Warn(feedback)
		utilities.WriteResponse(http.StatusUnauthorized, feedback, w)
		return
	}

	// test if the new username is already in the db
	_, errNewUsername := database.DBcon.GetIdByName(newUsername)
	if errNewUsername == nil {
		message := fmt.Sprintf("WARNING, the username %s is already taken, please choose another one", newUsername)
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusConflict, message, w)
		return
	}

	userid, errOldUsername := database.DBcon.GetIdByName(oldUsername)
	if errOldUsername != nil {
		logrus.Warn("Error while getting the user id from the client request")
		utilities.WriteResponse(http.StatusInternalServerError, "Error while getting the user id from the client request", w)
		return
	}

	errModifyName := database.DBcon.ModifyUsername(userid, newUsername)
	if errModifyName != nil {
		logrus.Warn(errModifyName.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errModifyName.Error(), w)
		return
	}

	utilities.WriteResponse(http.StatusOK, "Username successfully updated", w)
	logrus.Info("Update done!!")
}
