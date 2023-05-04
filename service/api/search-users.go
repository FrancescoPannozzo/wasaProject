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
func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Println("Searching users..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		logrus.Warn(utilities.Unauthorized)
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	targetUser := r.URL.Query().Get("username")

	errCheckUser := utilities.CheckUsername(targetUser)
	if errCheckUser != nil {
		logrus.Warn(errCheckUser.Error())
		utilities.WriteResponse(http.StatusBadRequest, errCheckUser.Error(), w)
		return
	}

	usernames, errGetNames := database.DBcon.GetUsernames(targetUser)

	if errGetNames != nil {
		logrus.Warn(errGetNames.Error())
		utilities.WriteResponse(http.StatusInternalServerError, errGetNames.Error(), w)
		return
	}

	if len(usernames) == 0 {
		utilities.WriteResponse(http.StatusNotFound, "User/s not found", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errEnc := json.NewEncoder(w).Encode(&usernames)
	if errEnc != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errEnc.Error(), w)
		return
	}
	logrus.Println("Done!")
}
