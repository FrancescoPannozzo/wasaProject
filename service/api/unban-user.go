package api

import (
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Unban the user provided by the path parameter username
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Removing the ban for the provided user..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		logrus.Warn(errId.Error())
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	errUsername := utilities.CheckUsername(ps.ByName("username"))
	if errUsername != nil {
		logrus.Warn("Bad request for the username to unban")
		utilities.WriteResponse(http.StatusBadRequest, "Username to unban not valid", w)
		return
	}

	if !database.DBcon.UsernameInDB(ps.ByName("username")) {
		message := "username not found"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	loggedUser, errNameId := rt.db.GetNameByID(utilities.GetBearerID(r))
	if errNameId != nil {
		message := "Unauthorized user"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusUnauthorized, message, w)
		return
	}

	feedback, err := database.DBcon.UnbanUser(loggedUser, ps.ByName("username"))
	if errors.Is(err, &utilities.DbBadRequest{}) {
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
	return
}
