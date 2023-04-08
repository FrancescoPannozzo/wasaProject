package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get an user profile
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	//prendo lista following
	// getfollowedlist()

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBaererID(r))

	//var thumbnails []Thumbnail

	thumbnails, err, httpStatus := database.DBcon.GetFollowedThumbnails(loggedUser)

	if err != nil {
		utilities.WriteResponse(httpStatus, err.Error(), w)
		return
	}

	result, err := json.Marshal(thumbnails)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
