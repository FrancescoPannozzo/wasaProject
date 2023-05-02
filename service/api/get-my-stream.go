package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get a user stream
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Getting the user stream..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		logrus.Warn(errId.Error())
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	// error not managed because GeNameById is already called in VerifyUserId
	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	thumbnails, errThumb := database.DBcon.GetFollowedThumbnails(loggedUser)
	if errThumb != nil {
		logrus.Warn(errThumb.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errThumb.Error(), w)
		return
	}

	result, errConv := json.Marshal(thumbnails)
	if errConv != nil {
		logrus.Warn(errConv.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errConv.Error(), w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	logrus.Infoln("Done!")
	return
}
