package api

import (
	"encoding/base64"
	"encoding/json"
	"fantastic-coffee-decaffeinated/service/database"
	"fantastic-coffee-decaffeinated/service/utilities"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Upload a user photo
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info("Uploading the photo..")

	errId := database.VerifyUserId(r, ps)
	if errId != nil {
		utilities.WriteResponse(http.StatusUnauthorized, errId.Error(), w)
		return
	}

	reqBody, errBody := io.ReadAll(r.Body)
	_ = r.Body.Close()

	if errBody != nil {
		message := "error while reading the body request"
		rt.baseLogger.WithError(errBody).Warning(message)
		utilities.WriteResponse(http.StatusBadRequest, message, w)
		return
	}

	dec, errDec := base64.StdEncoding.DecodeString(string(reqBody))
	if errDec != nil {
		message := "Server Error while decoding the photo"
		rt.baseLogger.WithError(errDec).Warning(message)
		utilities.WriteResponse(http.StatusInternalServerError, message, w)
		return
	}

	userId := utilities.GetBearerID(r)

	idphoto := userId[:4] + utilities.GenerateTimestamp()
	fileName := idphoto + ".png"
	filePath := filepath.Join("storage", fileName)
	tmpfile, errCreate := os.Create(filePath)
	if errCreate != nil {
		message := "Error with the creation of the file to store"
		rt.baseLogger.WithError(errCreate).Warning(message)
		utilities.WriteResponse(http.StatusInternalServerError, message, w)
		return
	}

	defer tmpfile.Close()

	_, errWrite := tmpfile.Write(dec)

	if errWrite != nil {
		message := "Error while writing the file"
		rt.baseLogger.WithError(errWrite).Warning(message)
		utilities.WriteResponse(http.StatusInternalServerError, message, w)
		return
	}

	if errSync := tmpfile.Sync(); errSync != nil {
		utilities.WriteResponse(http.StatusInternalServerError, "Error while writing the file", w)
		return
	}

	username, errNameId := rt.db.GetNameByID(userId)
	if errNameId != nil {
		message := "error while storing the file"
		rt.baseLogger.WithError(errNameId).Warning(message)
		utilities.WriteResponse(http.StatusInternalServerError, message, w)
		return
	}

	//Adding the photo data into the DB
	feedback, errPhoto := database.DBcon.InsertPhoto(username, idphoto)
	if errPhoto != nil {
		rt.baseLogger.WithError(errNameId).Warning(feedback)
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
		rt.baseLogger.WithError(errJson).Warning("wrong idphoto JSON")
		utilities.WriteResponse(http.StatusInternalServerError, "cannot read the request", w)
		return
	}

	logrus.Info("Done!")
	return
}
