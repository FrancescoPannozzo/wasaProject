package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// unFollow a user.
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	loggedUser, _ := database.DBcon.GetNameByID(utilities.GetBearerID(r))

	if ps.ByName("username") == loggedUser {
		utilities.WriteResponse(http.StatusBadRequest, "Warning, logged cannot unfollow himself", w)
	}

	feedback, err := database.DBcon.DeleteFollowed(loggedUser, ps.ByName("username"))

	if err != nil {
		logrus.Errorln(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	return
}
