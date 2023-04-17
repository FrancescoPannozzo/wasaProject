package api

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Remove a like
func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	errId := database.VerifyUserId(w, r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	username, errUsername := database.DBcon.GetNameByID(utilities.GetBearerID(r))

	if errUsername != nil {
		rt.baseLogger.WithError(errUsername).Warning("Cannot find the user")
		utilities.WriteResponse(http.StatusBadRequest, "Cannot find the user", w)
		return
	}

	//check if the user is the comment owner

	feedback, err := database.DBcon.RemoveComment(username, ps.ByName("idPhoto"), ps.ByName("idComment"))

	if errors.Is(err, sql.ErrNoRows) {
		utilities.WriteResponse(http.StatusBadRequest, feedback, w)
		return
	}

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	return
}
