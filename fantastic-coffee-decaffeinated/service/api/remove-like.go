package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Remove a like
func (rt *_router) removeLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	username, errUsername := rt.db.GetNameByID(database.GetBaererID(r))

	if errUsername != nil {
		rt.baseLogger.WithError(errUsername).Warning("Cannot find the user")
		utilities.WriteResponse(http.StatusBadRequest, "Cannot find the user", w)
		return
	}

	feedback, err, httpStatus := database.DBcon.RemoveLike(username, ps.ByName("idPhoto"))

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
	}

	utilities.WriteResponse(httpStatus, feedback, w)
	return
}
