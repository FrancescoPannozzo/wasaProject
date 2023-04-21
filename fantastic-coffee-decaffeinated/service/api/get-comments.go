package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a comment list of a user photo
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := database.VerifyUserId(r, ps)

	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	var comments []utilities.Comment

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	idPhoto := ps.ByName("idPhoto")
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

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("photo not found")
		utilities.WriteResponse(http.StatusBadRequest, "photo not found", w)
		return
	}

	// Check if the requested photo is in the DB
	_, errPhoto = database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errPhoto != nil {
		// Comment list retrive attempt of one photo not in the DB, so comments doesn' exists
		logrus.Warn(errPhoto.Error())
		utilities.WriteResponse(http.StatusNotFound, errPhoto.Error(), w)
		return
	}

	comments, errComm := database.DBcon.GetComments(loggedUser, ps.ByName("idPhoto"))
	if errComm != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errComm.Error(), w)
		return
	}

	result, errConv := json.Marshal(comments)
	if errConv != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errConv.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
