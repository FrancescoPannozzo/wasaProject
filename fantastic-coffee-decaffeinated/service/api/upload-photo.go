package api

import (
	"encoding/base64"
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info("Uploading photo")

	httpStatus, message := database.VerifyUseridController(w, r, ps)

	if httpStatus != http.StatusOK {
		utilities.WriteResponse(httpStatus, message, w)
		return
	}

	reqBody, errBody := io.ReadAll(r.Body)
	_ = r.Body.Close()

	if errBody != nil {
		err := fmt.Errorf("error while reading the body request: %v", errBody)
		logrus.Println(errBody)
		utilities.WriteResponse(http.StatusBadRequest, err.Error(), w)
		return
	}

	dec, errDec := base64.StdEncoding.DecodeString(string(reqBody))
	if errDec != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Server Error while decoding the photo", w)
		return
	}

	userId := utilities.GetBearerID(r)

	idphoto := userId[:4] + utilities.GenerateTimestamp()
	fileName := idphoto + ".png"
	filePath := filepath.Join("storage", fileName)
	tmpfile, errCreate := os.Create(filePath)
	if errCreate != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writingcreating", w)
		return
	}

	defer tmpfile.Close()

	breaded, errWrite := tmpfile.Write(dec)

	if errWrite != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writing the file", w)
		return
	}

	fmt.Println("Bytes readed: ", breaded)

	if errSync := tmpfile.Sync(); errSync != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writing the file", w)
		return
	}

	username, err := rt.db.GetNameByID(userId)

	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, username, w)
		return
	}

	//Adding a DB record
	feedback, err := database.DBcon.InsertPhoto(username, idphoto)
	if err != nil {
		utilities.WriteResponse(http.StatusInternalServerError, feedback, w)
		return
	}

	type Idphoto struct {
		Idphoto string `json:"idphoto"`
	}

	photo := &Idphoto{Idphoto: idphoto}
	w.WriteHeader(http.StatusCreated)
	errJson := json.NewEncoder(w).Encode(&photo)
	if errJson != nil {
		rt.baseLogger.WithError(err).Warning("wrong idphoto JSON")
		utilities.WriteResponse(http.StatusInternalServerError, "cannot read the request", w)
		return
	}

	return
}
