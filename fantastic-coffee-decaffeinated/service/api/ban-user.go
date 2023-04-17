package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// ban an user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Banning the provided user..")
	err := database.VerifyUserId(w, r, ps)

	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	loggedUser, err := rt.db.GetNameByID(utilities.GetBearerID(r))

	if err != nil {
		utilities.WriteResponse(http.StatusNotFound, loggedUser, w)
		return
	}

	type Banned struct {
		Username string `json:"name"`
	}

	var banned Banned
	_ = json.NewDecoder(r.Body).Decode(&banned)

	if loggedUser == banned.Username {
		utilities.WriteResponse(http.StatusBadRequest, "Logged user cannot ban himself", w)
		return
	}

	//check if the user to ban is in the DB
	_, errId := database.DBcon.GetIdByName(banned.Username)
	if errId != nil {
		utilities.WriteResponse(http.StatusBadRequest, errId.Error(), w)
		return
	}

	feedback, err := database.DBcon.BanUser(loggedUser, banned.Username)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}
	logrus.Infoln("Done!")
	utilities.WriteResponse(http.StatusCreated, feedback, w)
	return
}
