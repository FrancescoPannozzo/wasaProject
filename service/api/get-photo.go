package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a user photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Getting the user photo..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		logrus.Warn(errId.Error())
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	idphoto := ps.ByName("idPhoto")

	if !utilities.IsPhotoIdValid(idphoto) {
		logrus.Warn("Invalid photo ID")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photo ID", w)
		return
	}

	loggedUser, errNameId := rt.db.GetNameByID(utilities.GetBearerID(r))
	if errNameId != nil {
		logrus.Warn(utilities.Unauthorized)
		utilities.WriteResponse(http.StatusUnauthorized, utilities.Unauthorized, w)
		return
	}

	targetUser, errPhoto := database.DBcon.GetNameFromPhotoId(idphoto)
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

	fileName := idphoto + ".png"
	filePath := filepath.Join("storage", fileName)

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, "cannot get the photo", w)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_, errWrite := w.Write(buf)
	if errWrite != nil {
		logrus.Warn(errWrite.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errWrite.Error(), w)
	}
	logrus.Infoln("Done!")
}
