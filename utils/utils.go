package utils

import (
	"net/http"

	"github.com/budhip/url-shortener/model"
)

func GetStatusCode(err error) int {
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrUrlNotPresent:
		return http.StatusBadRequest
	case model.ErrSlugAlreadyUse:
		return http.StatusConflict
	case model.ErrSlugNotMatch:
		return http.StatusUnprocessableEntity
	case model.ErrSlugNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func StatusCode(w http.ResponseWriter, err error, status int) {
	if err != nil {
		w.WriteHeader(status)
		return
	}
	w.WriteHeader(status)
}