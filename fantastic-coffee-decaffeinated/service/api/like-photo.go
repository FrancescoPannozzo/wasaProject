package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Follow a user.
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	username, errUsername := rt.db.GetNameByID(utilities.GetBearerID(r))

	if errUsername != nil {
		logrus.Errorln("Cannot find the user")
		utilities.WriteResponse(http.StatusUnauthorized, "Cannot find the user", w)
		return
	}

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
