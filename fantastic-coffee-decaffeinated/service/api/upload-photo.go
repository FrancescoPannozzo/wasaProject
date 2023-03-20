package api

import (
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info("Uploading photo")

	httpStatus, message := database.VerifyUseridController(w, r)

	if httpStatus == 400 {
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	//auth control

	/*
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/

	//params := httprouter.ParamsFromContext(r.Context())

	fmt.Fprintf(w, "URL parameters:, %s, %s\n", ps.ByName("username"), ps.ByName("idPhoto"))

	//fmt.Println(ps.ByName("username"))

	return
}
