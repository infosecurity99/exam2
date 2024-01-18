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
//create trip
func (h Handler) CreateTrip(w http.ResponseWriter, r *http.Request) {

}
//getbyidtrip
func (h Handler) GetTripByID(w http.ResponseWriter, r *http.Request) {

}
//getlist trip
func (h Handler) GetTripList(w http.ResponseWriter) {

}
//updatetrip
func (h Handler) UpdateTrip(w http.ResponseWriter, r *http.Request) {

}
//deleate trip
func (h Handler) DeleteTrip(w http.ResponseWriter, r *http.Request) {

}
