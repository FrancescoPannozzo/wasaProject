package api

import (
	"encoding/json"
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// ban a user
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Infoln("Banning the provided user..")
	errId := database.VerifyUserId(r, ps)
	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	type Banned struct {
		Username string `json:"name"`
	}

	var banned Banned
	_ = json.NewDecoder(r.Body).Decode(&banned)

	errUsername := utilities.CheckUsername(banned.Username)
	if errUsername != nil {
		message := "User ID not allowed"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	if loggedUser == banned.Username {
		utilities.WriteResponse(http.StatusBadRequest, "Logged user cannot ban himself", w)
		return
	}

	//check if the user to ban is in the DB
	if !database.DBcon.UsernameInDB(banned.Username) {
		message := "User to ban not found"
		logrus.Warn(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	feedback, err := database.DBcon.BanUser(loggedUser, banned.Username)
	if errors.Is(err, &utilities.DbBadRequest{}) {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusBadRequest, feedback, w)
		return
	}
	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}
	logrus.Infoln("Done!")
	utilities.WriteResponse(http.StatusCreated, feedback, w)
	return
}
