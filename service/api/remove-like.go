package api

import (
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Remove a like
func (rt *_router) removeLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Warn("Removing the like..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	errUser := utilities.CheckUsername(ps.ByName("username"))
	if errUser != nil {
		logrus.Warn("User ID not valid")
		utilities.WriteResponse(http.StatusBadRequest, "User ID not valid", w)
		return
	}
	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("Invalid photo ID")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photo ID", w)
		return
	}

	if !rt.db.CheckOwnership(utilities.GetBearerID(r), ps.ByName("username")) {
		message := "The logged user can't remove a like of other users"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusUnauthorized, message, w)
		return
	}

	feedback, err := database.DBcon.RemoveLike(ps.ByName("username"), ps.ByName("idPhoto"))
	if errors.Is(err, &utilities.DbBadRequestError{}) {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusNotFound, feedback, w)
		return
	}
	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	logrus.Warn("Done!")
}
