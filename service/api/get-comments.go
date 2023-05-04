package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a comment list of a user photo
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Getting the comments..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	var comments []utilities.Comment

	loggedUser, errNameId := rt.db.GetNameByID(utilities.GetBearerID(r))
	if errNameId != nil {
		logrus.Warn(utilities.Unauthorized)
		utilities.WriteResponse(http.StatusUnauthorized, utilities.Unauthorized, w)
		return
	}

	idPhoto := ps.ByName("idPhoto")
	if !utilities.IsPhotoIdValid(idPhoto) {
		logrus.Warn("Invalid photo ID")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photo ID", w)
		return
	}
	targetUser, errPhoto := database.DBcon.GetNameFromPhotoId(idPhoto)
	if errPhoto != nil {
		logrus.Warn(errPhoto.Error())
		utilities.WriteResponse(http.StatusNotFound, errPhoto.Error(), w)
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
		logrus.Warn(errPhoto.Error())
		utilities.WriteResponse(http.StatusBadRequest, errPhoto.Error(), w)
		return
	}

	comments, errComm := database.DBcon.GetComments(loggedUser, ps.ByName("idPhoto"))
	if errors.Is(errComm, sql.ErrNoRows) {
		message := "Comments not found"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusNotFound, message, w)
		return
	}
	if errComm != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errComm.Error(), w)
		return
	}

	if comments == nil {
		logrus.Warn("No comments found")
		utilities.WriteResponse(http.StatusNotFound, "No comments found", w)
		return
	}

	result, errConv := json.Marshal(comments)
	if errConv != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errConv.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write(result)
	if errWrite != nil {
		logrus.Warn(errWrite.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errWrite.Error(), w)
	}

	logrus.Infoln("Done!")
}
