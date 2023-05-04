package api

import (
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// unFollow a user.
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Removing the follow..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	errUser := utilities.CheckUsername(ps.ByName("username"))
	if errUser != nil {
		message := "username to unfollow not valid"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	if !(database.DBcon.UsernameInDB(ps.ByName("username"))) {
		utilities.WriteResponse(http.StatusNotFound, "Warning, the user to unfollow is not in the DB", w)
		return
	}

	loggedUser, errNameId := database.DBcon.GetNameByID(utilities.GetBearerID(r))
	if errNameId != nil {
		logrus.Warn(utilities.Unauthorized)
		utilities.WriteResponse(http.StatusUnauthorized, utilities.Unauthorized, w)
		return
	}

	if ps.ByName("username") == loggedUser {
		utilities.WriteResponse(http.StatusBadRequest, "Warning, logged cannot unfollow himself", w)
	}

	feedback, err := database.DBcon.DeleteFollowed(loggedUser, ps.ByName("username"))
	if errors.Is(err, &utilities.DbBadRequestError{}) {
		logrus.Warn(feedback)
		utilities.WriteResponse(http.StatusNotFound, feedback, w)
		return
	}
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	logrus.Infoln("Done!")
}
