package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ban an user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	loggedUser, err := rt.db.GetNameByID(utilities.GetBearerID(r))
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, loggedUser, w)
		return
	}
	feedback, err := database.DBcon.UnbanUser(loggedUser, ps.ByName("username"))

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}
	utilities.WriteResponse(http.StatusOK, feedback, w)
	return

}
