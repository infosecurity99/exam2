package handler

import (
	"net/http"
)

func (h Handler) Customer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetCustomerList(w, r)
		} else {
			h.GetCustomerByID(w, r)
		}
	case http.MethodPut:
		h.UpdateCustomer(w, r)
	case http.MethodDelete:
		h.DeleteCustomer(w, r)
	}
}

func (h Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCustomerList(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

}
