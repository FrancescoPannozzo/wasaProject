package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Remove a like
func (rt *_router) removeLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	username, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	if !rt.db.CheckOwnership(utilities.GetBearerID(r), username) {
		utilities.WriteResponse(http.StatusUnauthorized, "The logged user can't remove a like of other users", w)
		return
	}

	errUser := utilities.CheckUsername(ps.ByName("username"))

	if errUser != nil || !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("User ID and/or photo id not valid")
		utilities.WriteResponse(http.StatusBadRequest, "User ID and/or photo id not valid", w)
		return
	}

	_, errPhoto := rt.db.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errPhoto != nil {
		logrus.Warn("photoId not found")
		utilities.WriteResponse(http.StatusNotFound, "photo not found", w)
		return
	}

	feedback, err := database.DBcon.RemoveLike(ps.ByName("username"), ps.ByName("idPhoto"))
	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	return
}
