package handler

import (
	"net/http"
)

func (h Handler) Driver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateDriver(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetDriverList(w, r)
		} else {
			h.GetDriverByID(w, r)
		}
	case http.MethodPut:
		h.UpdateDriver(w, r)
	case http.MethodDelete:
		h.DeleteDriver(w, r)
	}
}

func (h Handler) CreateDriver(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetDriverByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetDriverList(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) UpdateDriver(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteDriver(w http.ResponseWriter, r *http.Request) {

}
