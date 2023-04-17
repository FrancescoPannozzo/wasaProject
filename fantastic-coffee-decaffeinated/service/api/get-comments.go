package api

import (
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Get an user profile
func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := database.VerifyUserId(w, r, ps)

	if err != nil {
		utilities.WriteResponse(http.StatusUnauthorized, err.Error(), w)
		return
	}

	var comments []utilities.Comment
	// get comments

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	comments, errComm := database.DBcon.GetComments(loggedUser, ps.ByName("idPhoto"))

	if errComm != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errComm.Error(), w)
		return
	}

	result, errConv := json.Marshal(comments)
	if errConv != nil {
		utilities.WriteResponse(http.StatusInternalServerError, errConv.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
