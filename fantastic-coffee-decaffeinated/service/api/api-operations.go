package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"bytes"
	"fmt"
	"io"
	"log"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	check(err)
	fmt.Println("reqBody content is:/n", bytes.NewBuffer(reqBody).String())

	_, _ = w.Write(reqBody)
}