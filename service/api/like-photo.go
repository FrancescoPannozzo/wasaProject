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
	errId := database.VerifyUserId(r, ps)
	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("photo id not valid")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photoID", w)
		return
	}

	loggedUser, errNameId := rt.db.GetNameByID(utilities.GetBearerID(r))
	if errNameId != nil {
		message := "Unauthorized user"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusUnauthorized, message, w)
		return
	}

	targetUser, errPhoto := rt.db.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errPhoto != nil {
		message := "Photo id to like not found"
		rt.baseLogger.WithError(errPhoto).Warning(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		logrus.Warn("Banned user found")
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	feedback, err := database.DBcon.LikePhoto(loggedUser, ps.ByName("idPhoto"))
	if errors.Is(err, &utilities.DbBadRequest{}) {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusConflict, feedback, w)
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
