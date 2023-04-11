package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// unFollow a user.
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	name, errUser := database.DBcon.GetNameByID(utilities.GetBaererID(r))

	if errUser != nil {
		rt.baseLogger.WithError(errUser).Warning("error JSON format")
		utilities.WriteResponse(http.StatusInternalServerError, "error JSON format", w)
	}

	if !(database.DBcon.UsernameInDB(name)) {
		utilities.WriteResponse(http.StatusBadRequest, "Warning, the user is not in the DB", w)
		return
	}

	if ps.ByName("username") == name {
		utilities.WriteResponse(http.StatusBadRequest, "Warning, you cannot unfollow yourself", w)
	}

	feedback, err, httpStatus := database.DBcon.DeleteFollowed(name, ps.ByName("username"))

	if err != nil {
		logrus.Errorln(err.Error())
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
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
