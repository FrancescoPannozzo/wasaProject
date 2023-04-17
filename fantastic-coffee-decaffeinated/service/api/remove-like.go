package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Remove a like
func (rt *_router) removeLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	errId := database.VerifyUserId(w, r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	username, errUsername := rt.db.GetNameByID(utilities.GetBearerID(r))

	if !rt.db.CheckOwnership(utilities.GetBearerID(r), username) {
		utilities.WriteResponse(http.StatusUnauthorized, "Thelogged user can't set off a like of other users", w)
		return
	}

	if errUsername != nil {
		rt.baseLogger.WithError(errUsername).Warning("Cannot find the user")
		utilities.WriteResponse(http.StatusBadRequest, "Cannot find the user", w)
		return
	}

	feedback, err := database.DBcon.RemoveLike(ps.ByName("username"), ps.ByName("idPhoto"))

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	return
}
