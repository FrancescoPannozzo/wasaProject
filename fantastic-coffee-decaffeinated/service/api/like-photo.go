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
	err := database.VerifyUserId(r, ps)

	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	// GetNameById is called in VerifyUserId,the error is already managed, no needs to do the same here
	username, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("photo not found")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photoID", w)
		return
	}

	_, errPhoto := rt.db.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errPhoto != nil {
		message := "Photo id to like not found"
		logrus.Warn("")
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	feedback, err := database.DBcon.LikePhoto(username, ps.ByName("idPhoto"))
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
	return
}
