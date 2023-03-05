package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	statusNumber, payloadMessage := utilities.VerifyUseridController(w, r)

	if statusNumber == 400 {
		utilities.WriteResponse(http.StatusBadRequest, payloadMessage, w)
		return
	}

	oldUsername := r.URL.Query().Get("username")

	newUsername, errName := utilities.GetNameFromReq(r)

	if errName != nil {
		fmt.Println(errName)
		return
	}

	userid, _ := database.DBcon.GetIdByName(oldUsername)

	err := database.DBcon.ModifyUsername(userid, newUsername)

	if err != nil {
		fmt.Println(err)
		utilities.WriteResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

	utilities.WriteResponse(http.StatusCreated, "Username successfully updated", w)
	return

	//verifica l'auth, verificandol'auth ottengo token, estrapola nuovo username dalla
	// request, cerca record con il token e modifica lo username ad esso associato
}
