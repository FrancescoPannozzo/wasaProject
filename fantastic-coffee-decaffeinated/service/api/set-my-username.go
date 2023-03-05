package api

import (
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username, errName := utilities.GetUserFromReq(r)

	if errName != nil {
		fmt.Println(errName)
		return
	}

	//verifica l'auth, verificandol'auth ottengo token, estrapola nuovo username dalla
	// request, cerca record con il token e modifica lo username ad esso associato
}
