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
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Getting the user stream..")
	err := database.VerifyUserId(r, ps)

	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	// error not managed because GeNameById is already called in VerifyUserId
	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	thumbnails, err := database.DBcon.GetFollowedThumbnails(loggedUser)
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	result, err := json.Marshal(thumbnails)
	if err != nil {
		logrus.Warn(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	logrus.Infoln("Done!")
	return
}
