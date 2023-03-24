package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == 400 {
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	username := ps.ByName("username")
	idphoto := ps.ByName("idPhoto")

	feedback, err, httpStatus := database.DBcon.DeletePhoto(username, idphoto)

	if httpStatus == http.StatusInternalServerError {
		utilities.WriteResponse(httpStatus, feedback+".Error:"+err.Error(), w)
		return
	}

	utilities.WriteResponse(httpStatus, feedback, w)
	return

}
