package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/budhip/url-shortener/domain"
	"github.com/budhip/url-shortener/model"
	"github.com/budhip/url-shortener/utils"
	"github.com/gorilla/mux"
)

// UrlShortenerHandler  represent the http handler for url shortener
type UrlShortenerHandler struct {
	UrlShortUseCase domain.UrlShortenerUseCase
}

// NewUrlShortenerHandler will initialize the shortener/ resources endpoint
func NewUrlShortenerHandler(r *mux.Router, usu domain.UrlShortenerUseCase) {
	handler := &UrlShortenerHandler{
		UrlShortUseCase: usu,
	}
	r.HandleFunc("/shorten", handler.UrlShorten).Methods("POST")
	r.HandleFunc("/{shortcode:[a-zA-Z0-9]+}", handler.GetShortCode).Methods("GET")
}

// UrlShorten will store the url shorten by given request body
func (ush *UrlShortenerHandler) UrlShorten(w http.ResponseWriter, r *http.Request) {
	var shortenReq model.ShortenReq

	err := json.NewDecoder(r.Body).Decode(&shortenReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusCreated

	resp, errResp := ush.UrlShortUseCase.StoreShorten(ctx, shortenReq.URL, shortenReq.Slug)
	if errResp != nil {
		respErrResp := model.ResponseError{
			Message: errResp.Error(),
		}

		httpStatus = utils.GetStatusCode(errResp)

		marshallingResponse(w, respErrResp, errResp, httpStatus)

		return
	}

	marshallingResponse(w, resp, nil, httpStatus)
}

func marshallingResponse(w http.ResponseWriter, data interface{}, err error, statusCode int) {
	payload, errPayload := json.Marshal(data)
	if errPayload != nil {
		log.Println("err when marshalling payload")
	}

	utils.StatusCode(w, err, statusCode)
	_, errWrite := w.Write(payload)
	if errWrite != nil {
		log.Println("err when write payload")
	}
}

func (ush *UrlShortenerHandler) GetShortCode(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortcode"]

	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	httpStatus := http.StatusFound

	url, err := ush.UrlShortUseCase.GetShortCode(ctx, shortCode)
	if err != nil {
		respErrResp := model.ResponseError{
			Message: err.Error(),
		}

		httpStatus = utils.GetStatusCode(err)

		marshallingResponse(w, respErrResp, err, httpStatus)

		return
	}

	marshallingResponse(w, &model.GetShortCodeResp{Location: url}, nil, httpStatus)
}