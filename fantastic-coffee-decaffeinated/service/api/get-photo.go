package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

// Get an user photo
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	loggedUser, _ := rt.db.GetNameByID(utilities.GetBaererID(r))
	targetUser, _ := database.DBcon.GetNameFromPhotoId(ps.ByName("idPhoto"))

	// check if the user is banned
	if database.DBcon.CheckBan(loggedUser, targetUser) {
		utilities.WriteResponse(http.StatusUnauthorized, "the logged user is banned for the specific request", w)
		return
	}

	idphoto := ps.ByName("idPhoto")
	fileName := idphoto + ".png"
	filePath := filepath.Join("storage", fileName)

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "cannot get the photo", w)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
}
