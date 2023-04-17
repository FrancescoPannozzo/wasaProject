package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Get an user profile
func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	errId := database.VerifyUserId(w, r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	targetUser := r.URL.Query().Get("username")

	logrus.Println("Query param =", targetUser)

	err := utilities.CheckUsername(targetUser)
	if err != nil {
		utilities.WriteResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

	usernames, err := database.DBcon.GetUsernames(targetUser)

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	if len(usernames) == 0 {
		//username := utilities.Username{Name: "no users found"}
		//usernames = append(usernames, username)
		utilities.WriteResponse(http.StatusNotFound, "User/s not found", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&usernames)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

}
