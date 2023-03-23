package api

import (
	"encoding/base64"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Here /upload-photo endpoint")
	logrus.Info("Uploading photo")

	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus == 400 {
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	logrus.Infof("username is:%s", ps.ByName("username"))

	fmt.Fprintf(w, "URL parameters: %s\n", ps.ByName("username"))
	//fmt.Println(ps.ByName("username"))

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		errBody := fmt.Errorf("error while reading the body request: %v", err)
		logrus.Println(errBody)
		utilities.WriteResponse(http.StatusBadRequest, "error while reading the body request", w)
		return
	}

	dec, err := base64.StdEncoding.DecodeString(string(reqBody))
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Server Error while decoding the photo", w)
		return
	}

	//fmt.Println(dec)

	//generating the file name
	userId, _ := database.DBcon.GetIdByName(ps.ByName("username"))

	fileName := userId[:4] + utilities.GenerateTimestamp() + ".png"

	filePath := "storage/" + fileName

	tmpfile, err := os.Create(filePath)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writingcreating", w)
		return
	}

	defer tmpfile.Close()

	breaded, err := tmpfile.Write(dec)

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writing the file", w)
		return
	}

	fmt.Println("Bytes readed: ", breaded)

	if err := tmpfile.Sync(); err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writing the file", w)
		return
	}

	utilities.WriteResponse(http.StatusCreated, "Photo uploaded", w)
	return
}
