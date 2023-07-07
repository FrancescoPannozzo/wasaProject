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
	"strings"

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

	base64data := string(reqBody)

	data := strings.Split(base64data, ",")
	var extension string

	switch data[0] {
	case "data:image/jpeg;base64":
		extension = ".jpeg"
		break
	case "data:image/png;base64":
		extension = ".png"
		break
	case "data:image/jpg;base64":
		extension = ".jpg"
		break
	default:
		extension = "unknown"
	}

	if extension == "unknown" {
		message := "Warning the photo extension is not supported, please load .jpeg, .png or .jpg files"
		utilities.WriteResponse(http.StatusUnsupportedMediaType, message, w)
		return
	}

	dec, errDec := base64.StdEncoding.DecodeString(data[1])
	if errDec != nil {
		message := "Server Error while decoding the photo"
		rt.baseLogger.WithError(errDec).Warning(message)
		utilities.WriteResponse(http.StatusInternalServerError, message, w)
		return
	}

	userId := utilities.GetBearerID(r)

	idphoto := userId[:4] + utilities.GenerateTimestamp()
	fileName := idphoto + extension

	pathStorage := "/tmp/media"

	if errDir := os.MkdirAll(pathStorage, os.ModePerm); errDir != nil {
		logrus.Error("Can't create a media storage directory for the photo storage")
		utilities.WriteResponse(http.StatusInternalServerError, "error with the storing path of the photo", w)
		return
	}

	filePath := filepath.Join(pathStorage, fileName)

	tmpfile, errCreate := os.Create(filePath)
	if errCreate != nil {
		message := "Error with the creation of the file to store, path:" + filePath
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

	// Adding the photo data into the DB
	feedback, errPhoto := database.DBcon.InsertPhoto(username, idphoto, extension)
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
}
