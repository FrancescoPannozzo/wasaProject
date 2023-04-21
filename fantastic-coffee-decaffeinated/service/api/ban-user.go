package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// ban a user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Banning the provided user..")
	err := database.VerifyUserId(r, ps)
	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

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
		utilities.WriteResponse(http.StatusNotFound, errId.Error(), w)
		return
	}

	feedback, err := database.DBcon.BanUser(loggedUser, banned.Username)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}
	logrus.Infoln("Done!")
	utilities.WriteResponse(http.StatusOK, feedback, w)
	return
}
