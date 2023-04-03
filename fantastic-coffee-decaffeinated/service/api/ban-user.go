package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ban an user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

}
