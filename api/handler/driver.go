package handler

import (
	"encoding/json"
	"errors"
	"exam2/api/models"
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

// create  drivers
func (h Handler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	createDrivers := models.CreateDriver{}

	if err := json.NewDecoder(r.Body).Decode(&createDrivers); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Driver().Create(createDrivers)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	customer, err := h.storage.Driver().Get(models.PrimaryKey{
		ID: pKey,
	}.ID)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, customer)
}

//get by id   drivers
func (h Handler) GetDriverByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	drivers, err := h.storage.Driver().Get(models.PrimaryKey{
		ID: id,
	}.ID)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, drivers)
}

//get  list drivers
func (h Handler) GetDriverList(w http.ResponseWriter, r *http.Request) {

}

//updated drivers
func (h Handler) UpdateDriver(w http.ResponseWriter, r *http.Request) {

}

//delete drivers
func (h Handler) DeleteDriver(w http.ResponseWriter, r *http.Request) {

}
