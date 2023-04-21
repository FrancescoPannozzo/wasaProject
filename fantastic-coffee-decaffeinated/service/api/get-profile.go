package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a user profile
func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := database.VerifyUserId(r, ps)

	if err != nil {
		logrus.Warn("Unauthorized user")
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	targetUser := ps.ByName("username")

	errUsername := utilities.CheckUsername(targetUser)
	if errUsername != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusBadRequest, errUsername.Error(), w)
		return
	}

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	thumbnails, err := database.DBcon.GetThumbnails(targetUser)

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&thumbnails)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

}
