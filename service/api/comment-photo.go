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
	logrus.Infoln("Posting the comment..")
	errId := database.VerifyUserId(r, ps)

	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))
	targetUser, errID := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))
	if errID != nil {
		utilities.WriteResponse(http.StatusBadRequest, "The photo id provided is not in the DB", w)
		return
	}

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		logrus.Warn("Banned user found")
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
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

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("Photo id not valid")
		utilities.WriteResponse(http.StatusBadRequest, "Photo id not valid", w)
	}

	feedback, err := database.DBcon.CommentPhoto(loggedUser, ps.ByName("idPhoto"), content.Comment)
	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
	}

	utilities.WriteResponse(http.StatusCreated, feedback, w)
	logrus.Infoln("Done!")
	return

}
