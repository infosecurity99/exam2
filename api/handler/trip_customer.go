package handler

import (
	"net/http"
)

func (h Handler) TripCustomer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTripCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetTripCustomerList(w)
		} else {
			h.GetTripCustomerByID(w, r)
		}
	case http.MethodPut:
		h.UpdateTripCustomer(w, r)
	case http.MethodDelete:
		h.DeleteTripCustomer(w, r)
	}
}

func (h Handler) CreateTripCustomer(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetTripCustomerByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetTripCustomerList(w http.ResponseWriter) {

}

func (h Handler) UpdateTripCustomer(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteTripCustomer(w http.ResponseWriter, r *http.Request) {

}
