package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Follow a user.
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	//name is the user to follow
	name, errReq := utilities.GetNameFromReq(r)

	if errReq != nil {
		rt.baseLogger.WithError(errReq).Warning("error JSON format")
		utilities.WriteResponse(http.StatusInternalServerError, "error JSON format", w)
	}

	if !(database.DBcon.UsernameInDB(name)) {
		utilities.WriteResponse(http.StatusBadRequest, "Warning, the user is not in the DB", w)
		return
	}

	if ps.ByName("username") == name {
		utilities.WriteResponse(http.StatusBadRequest, "Warning, you cannot follow yourself", w)
	}

	//Quindi inserisci il follower come tale nel DB
	feedback, err, httpStatus := database.DBcon.InsertFollower(ps.ByName("username"), name)

	if err != nil {
		logrus.Errorln(err.Error())
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	return
	/*
		err := json.NewDecoder(r.Body).Decode(&username)
		_ = r.Body.Close()
		if err != nil {
			rt.baseLogger.Warning()
			utilities.WriteResponse(http.StatusInternalServerError, "error JSON format", w)
		}
	*/
}
