package handler

import (
	"net/http"
)

func (h Handler) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCar(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetCarList(w)
		} else {
			h.GetCarByID(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		if _, ok := values["route"]; ok {
			h.UpdateCarRoute(w, r)
		} else if _, ok := values["status"]; ok {
			h.UpdateCarStatus(w, r)
		} else {
			h.UpdateCar(w, r)
		}
	case http.MethodDelete:
		h.DeleteCar(w, r)
	}
}

func (h Handler) CreateCar(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCarByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCarList(w http.ResponseWriter) {

}

func (h Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) UpdateCarRoute(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) UpdateCarStatus(w http.ResponseWriter, r *http.Request) {

}
