package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Delete an user photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	// checking if the logged user has the rights to perform the action
	loggedUsername, err := rt.db.GetNameByID(utilities.GetBaererID(r))
	if err != nil {
		utilities.WriteResponse(http.StatusNotFound, err.Error(), w)
	}
	photoOwner, err := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "cannot process the request", w)
		return
	}
	if loggedUsername != photoOwner {
		utilities.WriteResponse(http.StatusUnauthorized, "User doesn't have the rights to delete the photo, operation refused", w)
		return
	}

	//loggedUsername := ps.ByName("loggedUsername")

	idphoto := ps.ByName("idPhoto")

	feedback, err, httpStatus := database.DBcon.DeletePhoto(idphoto)

	if httpStatus == http.StatusInternalServerError {
		utilities.WriteResponse(httpStatus, feedback+".Error:"+err.Error(), w)
		return
	}

	utilities.WriteResponse(httpStatus, feedback, w)
	return

}
