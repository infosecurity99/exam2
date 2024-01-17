package handler

import (
	"net/http"
)

func (h Handler) City(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCity(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetCityList(w)
		} else {
			h.GetCityByID(w, r)
		}
	case http.MethodPut:
		h.UpdateCity(w, r)
	case http.MethodDelete:
		h.DeleteCity(w, r)
	}
}

func (h Handler) CreateCity(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCityByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCityList(w http.ResponseWriter) {

}

func (h Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteCity(w http.ResponseWriter, r *http.Request) {

}
