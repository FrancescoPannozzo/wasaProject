package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Remove a like
func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	username, errUsername := database.DBcon.GetNameByID(utilities.GetBaererID(r))

	if errUsername != nil {
		rt.baseLogger.WithError(errUsername).Warning("Cannot find the user")
		utilities.WriteResponse(http.StatusBadRequest, "Cannot find the user", w)
		return
	}

	//check if the user is the comment owner

	feedback, err, httpStatus := database.DBcon.RemoveComment(username, ps.ByName("idPhoto"), ps.ByName("idComment"))

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
	}

	utilities.WriteResponse(httpStatus, feedback, w)
	return
}
