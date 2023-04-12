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
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	//loggedUser, _ := rt.db.GetNameByID(utilities.GetBaererID(r))
	targetUser := r.URL.Query().Get("username")

	logrus.Println("Query param =", targetUser)

	err := utilities.CheckUsername(targetUser)
	if err != nil {
		utilities.WriteResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

	usernames, err, httpStatus := database.DBcon.GetUsernames(targetUser)

	if err != nil {
		utilities.WriteResponse(httpStatus, err.Error(), w)
		return
	}

	if len(usernames) == 0 {
		username := utilities.Username{Name: "no users found"}
		usernames = append(usernames, username)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&usernames)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

}
