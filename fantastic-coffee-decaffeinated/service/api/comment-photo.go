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
	errId := database.VerifyUserId(w, r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	username, errUsername := rt.db.GetNameByID(utilities.GetBearerID(r))

	if errUsername != nil {
		logrus.Errorln("Cannot find the user")
		utilities.WriteResponse(http.StatusNotFound, "Cannot find the user", w)
		return
	}

	type Comment struct {
		Comment string `json:"comment"`
	}

	var content Comment
	_ = json.NewDecoder(r.Body).Decode(&content)

	if len(content.Comment) > 100 {
		utilities.WriteResponse(http.StatusBadRequest, "Comment provided is longer than 100 characters", w)
		return
	}

	_, errID := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errID != nil {
		utilities.WriteResponse(http.StatusBadRequest, "The photo id provided is not in the DB", w)
		return
	}

	feedback, err := database.DBcon.CommentPhoto(username, ps.ByName("idPhoto"), content.Comment)

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	return

}
