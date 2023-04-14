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
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	var comments []utilities.Comment
	// get comments

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBearerID(r))

	comments, err := database.DBcon.GetComments(loggedUser, ps.ByName("idPhoto"))

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	result, err := json.Marshal(comments)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
