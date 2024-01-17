package handler

import (
	"net/http"
)

func (h Handler) Trip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTrip(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetTripList(w)
		} else {
			h.GetTripByID(w, r)
		}
	case http.MethodPut:
		h.UpdateTrip(w, r)
	case http.MethodDelete:
		h.DeleteTrip(w, r)
	}
}

func (h Handler) CreateTrip(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetTripByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetTripList(w http.ResponseWriter) {

}

func (h Handler) UpdateTrip(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteTrip(w http.ResponseWriter, r *http.Request) {

}
