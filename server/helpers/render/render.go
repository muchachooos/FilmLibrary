package render

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, resp interface{}) {
	response, err := json.Marshal(resp)
	if err != nil {
		Internal(w, err)
		return
	}

	OK(w)

	_, err = w.Write(response)
	if err != nil {
		log.Printf("fail to write response: %s", err)
	}
}

func OK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func Internal(w http.ResponseWriter, err error) {
	log.Printf("Internal error: %s", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	writeErrorToResponse(w, err)
}

func BadRequest(w http.ResponseWriter, err error) {
	log.Printf("Bad Request error: %s", err.Error())
	w.WriteHeader(http.StatusBadRequest)
	writeErrorToResponse(w, err)
}

func NotFound(w http.ResponseWriter) {
	log.Print("Not Found error")
	w.WriteHeader(http.StatusNotFound)
}

func Unauthorized(w http.ResponseWriter) {
	log.Print("Unauthorized error")
	w.WriteHeader(http.StatusUnauthorized)
}

func NotAllowed(w http.ResponseWriter) {
	log.Print("Not Allowed error")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func writeErrorToResponse(w http.ResponseWriter, err error) {
	_, errW := w.Write([]byte(fmt.Sprintf("database error: %s", err)))
	if errW != nil {
		log.Printf("fail to write error to response: %s", errW)
	}
}
