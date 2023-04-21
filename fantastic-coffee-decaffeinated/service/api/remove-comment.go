package api

import (
	"database/sql"
	"errors"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Remove a comment
func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	errId := database.VerifyUserId(r, ps)
	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	if !utilities.IsPhotoIdValid(ps.ByName("idPhoto")) {
		logrus.Warn("photo id not valid")
		utilities.WriteResponse(http.StatusBadRequest, "photo id not valid", w)
		return
	}

	//check if the user is the comment owner
	loggedUser, _ := database.DBcon.GetNameByID(utilities.GetBearerID(r))
	commentOwner, errUser := database.DBcon.GetNameFromCommentId(ps.ByName("idComment"))
	if errUser != nil {
		logrus.Warn(errUser.Error())
		utilities.WriteResponse(http.StatusNotFound, "comment not found", w)
		return
	}

	if loggedUser != commentOwner {
		logrus.Warn("Unauthorized to perform this action")
		utilities.WriteResponse(http.StatusUnauthorized, "the user is not the owner of the comment", w)
		return
	}

	feedback, err := database.DBcon.RemoveComment(ps.ByName("idComment"))
	if errors.Is(err, sql.ErrNoRows) {
		utilities.WriteResponse(http.StatusNotFound, feedback, w)
		return
	}

	if err != nil {
		rt.baseLogger.WithError(err).Warning(feedback)
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	utilities.WriteResponse(http.StatusOK, feedback, w)
	return
}
