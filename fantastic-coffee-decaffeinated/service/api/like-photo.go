package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Follow a user.
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := database.VerifyUserId(w, r, ps)

	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	// GetNameById is called in VerifyUserId,the error is already managed, no needs to do the same here
	username, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	feedback, err := rt.db.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if err != nil {
		utilities.WriteResponse(http.StatusBadRequest, feedback, w)
		return
	}

	feedback, err = database.DBcon.LikePhoto(username, ps.ByName("idPhoto"))

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	return
}
