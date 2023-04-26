package api

import (
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Give a like to a user photo.
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Posting a like to the photo..")
	err := database.VerifyUserId(r, ps)

	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	// GetNameById is called in VerifyUserId,the error is already managed, no needs to do the same here
	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	targetUser, errPhoto := rt.db.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errPhoto != nil {
		message := "Photo id to like not found"
		logrus.Warn("")
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		logrus.Warn("Banned user found")
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("photo not found")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photoID", w)
		return
	}

	feedback, err := database.DBcon.LikePhoto(loggedUser, ps.ByName("idPhoto"))
	if errors.Is(err, &utilities.DbBadRequest{}) {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusBadRequest, feedback, w)
		return
	}
	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	logrus.Infoln("Done!")
	return
}
