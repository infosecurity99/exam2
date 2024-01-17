package handler

import (
	"encoding/json"
	"exam2/api/models"
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

//create  city
func (h Handler) CreateCity(w http.ResponseWriter, r *http.Request) {
	newCity := models.City{}

	if err := json.NewDecoder(r.Body).Decode(&newCity); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cityID, err := h.storage.City().Create(models.CreateCity{})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdCity, err := h.storage.City().Get(cityID)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusCreated, createdCity)
}



func (h Handler) GetCityByID(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) GetCityList(w http.ResponseWriter) {

}

func (h Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) DeleteCity(w http.ResponseWriter, r *http.Request) {

}
