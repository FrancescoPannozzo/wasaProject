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
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	username, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	type Comment struct {
		Comment string `json:"comment"`
	}

	var content Comment
	_ = json.NewDecoder(r.Body).Decode(&content)

	if len(content.Comment) > 100 {
		utilities.WriteResponse(http.StatusBadRequest, "Comment provided is longer than 100 characters", w)
		return
	}

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("Photo not found")
		utilities.WriteResponse(http.StatusBadRequest, "Photo not found", w)
	}

	_, errID := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errID != nil {
		utilities.WriteResponse(http.StatusNotFound, "The photo id provided is not in the DB", w)
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
