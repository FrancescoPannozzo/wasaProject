package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a user post
func (rt *_router) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := database.VerifyUserId(r, ps)

	if err != nil {
		logrus.Warn("Unauthorized user")
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	idPhoto := ps.ByName("idPhoto")
	if !utilities.IsPhotoIdValid(idPhoto) {
		logrus.Warn("photo id not valid")
		utilities.WriteResponse(http.StatusBadRequest, "photo id not valid", w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	targetUser, errPhoto := database.DBcon.GetNameFromPhotoId(idPhoto)
	if errPhoto != nil {
		logrus.Warn(errPhoto.Error())
		utilities.WriteResponse(http.StatusNotFound, err.Error(), w)
		return
	}

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		logrus.Warn("Banned user found")
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	post, err := database.DBcon.GetPost(loggedUser, idPhoto)
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
	}

	result, errConv := json.Marshal(post)
	if errConv != nil {
		logrus.Warn(errConv.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errConv.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	logrus.Infoln("Done!")
	return
}
