package handler

import (
	"encoding/json"
	"errors"
	"exam2/api/models"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCar(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetCarList(w, r)
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

//create car
func (h Handler) CreateCar(w http.ResponseWriter, r *http.Request) {
	newsCars := models.CreateCar{}

	if err := json.NewDecoder(r.Body).Decode(&newsCars); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	carId, err := h.storage.Car().Create(newsCars)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdCar, err := h.storage.Car().Get(models.PrimaryKey{ID: carId})
	fmt.Println(createdCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, createdCar)
}

//getbyid vcar
func (h Handler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	car, err := h.storage.Car().Get(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, car)
}

//getlist car
func (h Handler) GetCarList(w http.ResponseWriter, r *http.Request) {
	var (
		page, limit = 1, 10
		err         error
	)
	values := r.URL.Query()

	if len(values["page"]) > 0 {
		page, err = strconv.Atoi(values["page"][0])
		if err != nil {
			page = 1
		}
	}

	if len(values["limit"]) > 0 {
		limit, err = strconv.Atoi(values["limit"][0])
		if err != nil {
			fmt.Println("limit", values["limit"])
			limit = 10
		}
	}

	resp, err := h.storage.Car().GetList(models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

//updatecar
func (h Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	updateCar := models.City{}

	if err := json.NewDecoder(r.Body).Decode(&updateCar); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.City().Update(updateCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storage.Car().Get(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, user)
}

//delate car
func (h Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	if err := h.storage.Car().Delete(models.PrimaryKey{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "data successfully deleted")
}

//updatecarroute
func (h Handler) UpdateCarRoute(w http.ResponseWriter, r *http.Request) {
	var updateCarRoute models.UpdateCarRoute

	err := json.NewDecoder(r.Body).Decode(&updateCarRoute)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.storage.Car().UpdateCarRoute(updateCarRoute)
	if err != nil {
		http.Error(w, "Failed to update car route", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}


// UpdateCarStatus updates the status of a car
func (h Handler) UpdateCarStatus(w http.ResponseWriter, r *http.Request) {
	var updateCarStatus models.UpdateCarStatus

	err := json.NewDecoder(r.Body).Decode(&updateCarStatus)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.storage.Car().UpdateCarStatus(updateCarStatus)
	if err != nil {
		http.Error(w, "Failed to update car status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
