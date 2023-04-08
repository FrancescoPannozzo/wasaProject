package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get an user profile
func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBaererID(r))
	targetUser := ps.ByName("username")

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	//var thumbnails []Thumbnail

	thumbnails, err, httpStatus := database.DBcon.GetThumbnails(targetUser)

	if err != nil {
		utilities.WriteResponse(httpStatus, err.Error(), w)
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