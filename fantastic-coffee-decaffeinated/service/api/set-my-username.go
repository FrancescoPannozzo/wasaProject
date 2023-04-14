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
	httpStatus, payloadMessage := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		logrus.Errorln("Error with the authentication, httpStatus is '%v', %s", httpStatus, payloadMessage)
		utilities.WriteResponse(httpStatus, payloadMessage, w)
		return
	}
	//oldUsername := r.URL.Query().Get("username")
	// Check if who wants to change the username is the real profile owner
	userId := utilities.GetBearerID(r)
	oldUsername := ps.ByName("username")

	if !database.DBcon.CheckOwnership(userId, oldUsername) {
		utilities.WriteResponse(http.StatusUnauthorized, "attempt to change someone else's username detected", w)
		return
	}

	// Get the username to change from the RequestBody
	newUsername, errName := utilities.GetNameFromReq(r)
	if errName != nil {
		logrus.Errorln("Error in setMyUsername() while getting the username from the client request %v", errName)
		utilities.WriteResponse(http.StatusBadRequest, "Error: requestBody not valid", w)
		return
	}

	fmt.Println("New username: ", newUsername)

	err := utilities.CheckUsername(newUsername)
	if err != nil {
		utilities.WriteResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

	//test if the newusername is already in the db
	userid, errDb := database.DBcon.GetIdByName(newUsername)
	if errDb == nil {
		message := fmt.Sprintf("WARNING, the username %s is already taken, please choose another one", newUsername)
		logrus.Println(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	userid, errDb = database.DBcon.GetIdByName(oldUsername)
	if errDb != nil {
		logrus.Infof("Error in setMyUsername() while getting the user id from the client request %v", errName)
		utilities.WriteResponse(http.StatusInternalServerError, "Error while getting the user id from the client request", w)
		return
	}

	err = database.DBcon.ModifyUsername(userid, newUsername)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	utilities.WriteResponse(http.StatusCreated, "Username successfully updated", w)
	return
}
