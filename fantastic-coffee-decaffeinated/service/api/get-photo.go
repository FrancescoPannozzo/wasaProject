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
	err := database.VerifyUserId(r, ps)

	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	idphoto := ps.ByName("idPhoto")

	if !utilities.IsPhotoIdValid(idphoto) {
		logrus.Warn("Invalid photo ID")
		utilities.WriteResponse(http.StatusBadRequest, "Invalid photo ID", w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	targetUser, errPhoto := database.DBcon.GetNameFromPhotoId(idphoto)
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

	fileName := idphoto + ".png"
	filePath := filepath.Join("storage", fileName)

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, "cannot get the photo", w)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
	logrus.Infoln("Done!")
	return
}
