package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ban an user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	loggedUser, err := rt.db.GetNameByID(utilities.GetBaererID(r))

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, loggedUser, w)
		return
	}

	type Banned struct {
		Username string `json:"username"`
	}

	var banned Banned
	_ = json.NewDecoder(r.Body).Decode(&banned)

	feedback, err, httpStatus := database.DBcon.BanUser(loggedUser, banned.Username)
	utilities.WriteResponse(httpStatus, feedback, w)
	return
}
