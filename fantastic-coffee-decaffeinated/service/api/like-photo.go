package api

import (
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

	feedback, err := rt.db.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if err != nil {
		utilities.WriteResponse(http.StatusNotFound, feedback, w)
		return
	}

	feedback, err = database.DBcon.LikePhoto(username, ps.ByName("idPhoto"))

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	return
}
