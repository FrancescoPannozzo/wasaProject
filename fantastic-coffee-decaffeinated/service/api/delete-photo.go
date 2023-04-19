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
	logrus.Warn("Deleting the photo..")
	err := database.VerifyUserId(r, ps)
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	// checking if the logged user has the rights to perform the action
	loggedUsername, err := rt.db.GetNameByID(utilities.GetBearerID(r))
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusNotFound, err.Error(), w)
	}
	photoOwner, err := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, "cannot process the request", w)
		return
	}
	if loggedUsername != photoOwner {
		utilities.WriteResponse(http.StatusUnauthorized, "User doesn't have the rights to delete the photo, operation refused", w)
		return
	}

	idphoto := ps.ByName("idPhoto")
	feedback, err, httpStatus := database.DBcon.DeletePhoto(idphoto)

	if httpStatus == http.StatusInternalServerError {
		utilities.WriteResponse(httpStatus, feedback+".Error:"+err.Error(), w)
		return
	}

	utilities.WriteResponse(httpStatus, feedback, w)
	return

}
