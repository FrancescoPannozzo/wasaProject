package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Delete a user photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Deleting the photo..")
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

	// checking if the logged user has the rights to perform the action
	loggedUsername, errNameID := rt.db.GetNameByID(utilities.GetBearerID(r))
	if errNameID != nil {
		logrus.Warn(errNameID.Error())
		utilities.WriteResponse(http.StatusNotFound, errNameID.Error(), w)
	}
	photoOwner, err := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusNotFound, "cannot process the request", w)
		return
	}
	if loggedUsername != photoOwner {
		message := "Logged user doesn't have the rights to delete the photo, operation refused"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusUnauthorized, message, w)
		return
	}

	feedback, errDel := database.DBcon.DeletePhoto(idphoto)
	if errDel != nil {
		rt.baseLogger.WithError(errDel).Warning("Error while deleting the photo")
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	logrus.Infoln("Done!")
}
