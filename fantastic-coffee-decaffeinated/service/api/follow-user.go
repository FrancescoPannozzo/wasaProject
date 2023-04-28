package api

import (
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Follow a user.
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Following the user..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	//name is the user to follow
	userToFollow, _ := utilities.GetNameFromReq(r)

	errUser := utilities.CheckUsername(userToFollow)
	if errUser != nil {
		message := "username not valid"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	//check if the user to follow is in the DB
	if !(database.DBcon.UsernameInDB(userToFollow)) {
		message := "Warning, the user provided is not in the DB"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusUnprocessableEntity, message, w)
		return
	}

	//Check if the user is trying to follow himself
	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	if loggedUser == userToFollow {
		utilities.WriteResponse(http.StatusConflict, "Warning, you cannot follow yourself", w)
		return
	}

	//Insert the user to follow in the DB
	feedback, errIns := database.DBcon.InsertFollower(loggedUser, userToFollow)
	if errors.Is(errIns, &utilities.DbBadRequest{}) {
		logrus.Warn(feedback)
		utilities.WriteResponse(http.StatusConflict, feedback, w)
		return
	}

	if errIns != nil {
		logrus.Warn(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	logrus.Infoln("Done!")
	return
}
