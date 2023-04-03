package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Comment a photo
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == http.StatusBadRequest || httpStatus == http.StatusUnauthorized {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	username, errUsername := rt.db.GetNameByID(utilities.GetBaererID(r))

	if errUsername != nil {
		logrus.Errorln("Cannot find the user")
		utilities.WriteResponse(http.StatusUnauthorized, "Cannot find the user", w)
		return
	}

	type Comment struct {
		Comment string `json:"comment"`
	}

	var content Comment
	_ = json.NewDecoder(r.Body).Decode(&content)

	feedback, err, httpStatus := database.DBcon.CommentPhoto(username, ps.ByName("idPhoto"), content.Comment)

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
	}

	utilities.WriteResponse(httpStatus, feedback, w)

}
