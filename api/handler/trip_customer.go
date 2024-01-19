package handler

import (
	"encoding/json"
	"errors"
	"exam2/api/models"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) TripCustomer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTripCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetTripCustomerList(w, r)
		} else {
			h.GetTripCustomerByID(w, r)
		}
	case http.MethodPut:
		h.UpdateTripCustomer(w, r)
	case http.MethodDelete:
		h.DeleteTripCustomer(w, r)
	}
}

//create  trip customers
func (h Handler) CreateTripCustomer(w http.ResponseWriter, r *http.Request) {
	newsTripCustomer := models.CreateTripCustomer{}

	if err := json.NewDecoder(r.Body).Decode(&newsTripCustomer); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	carId, err := h.storage.TripCustomer().Create(newsTripCustomer)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdCar, err := h.storage.TripCustomer().Get((models.PrimaryKey{ID: carId}))
	fmt.Println(createdCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, createdCar)
}

//get by id trip customers
func (h Handler) GetTripCustomerByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	tripcustomer, err := h.storage.TripCustomer().Get(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, tripcustomer)
}

//get list trip customers
func (h Handler) GetTripCustomerList(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.storage.TripCustomer().GetList(models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

// update trip customers
func (h Handler) UpdateTripCustomer(w http.ResponseWriter, r *http.Request) {

	updateTripCustomers := models.TripCustomer{}

	if err := json.NewDecoder(r.Body).Decode(&updateTripCustomers); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.TripCustomer().Update(updateTripCustomers)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	tripcostumer, err := h.storage.TripCustomer().Get(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(tripcostumer)
	handleResponse(w, http.StatusOK, tripcostumer)
}

//delere   trip customers
func (h Handler) DeleteTripCustomer(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	if err := h.storage.TripCustomer().Delete(models.PrimaryKey{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "data successfully deleted")
}
