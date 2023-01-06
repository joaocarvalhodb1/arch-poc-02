package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Erro struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func JsonResponse(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println(err.Error())
		}
	}
}

func JsonError(w http.ResponseWriter, statusCode int, err error) {
	JsonResponse(w, statusCode, Erro{
		Code:  statusCode,
		Error: err.Error(),
	})
}

func Params(param string, w http.ResponseWriter, r *http.Request) (string, error) {
	params := chi.URLParam(r, param)
	if params == "" {
		err := errors.New("Error reading parameter " + param)
		JsonError(w, http.StatusBadRequest, err)
		return "", err
	}
	return params, nil
}
