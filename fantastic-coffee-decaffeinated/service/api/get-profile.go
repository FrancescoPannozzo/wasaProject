package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a user profile
func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Getting the user profile..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		logrus.Warn("Unauthorized user")
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	targetUser := ps.ByName("username")

	errUsername := utilities.CheckUsername(targetUser)
	if errUsername != nil {
		logrus.Warn(errUsername.Error())
		utilities.WriteResponse(http.StatusBadRequest, errUsername.Error(), w)
		return
	}

	if !database.DBcon.UsernameInDB(targetUser) {
		message := "user not found"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusNotFound, message, w)
		return
	}

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	thumbnails, errThumb := database.DBcon.GetThumbnails(targetUser)
	if errThumb != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errThumb.Error(), w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errEnc := json.NewEncoder(w).Encode(&thumbnails)
	if errEnc != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errEnc.Error(), w)
		return
	}
	logrus.Infoln("Done!")
	return
}
