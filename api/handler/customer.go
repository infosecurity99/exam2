package handler

import (
	"encoding/json"
	"errors"
	"exam2/api/models"
	"exam2/check"
	"fmt"
	"net/http"
	"strconv"
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

//createcustomers
func (h Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	createCustomers := models.CreateCustomer{}

	if err := json.NewDecoder(r.Body).Decode(&createCustomers); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}
	if !check.PhoneNumber(createCustomers.Phone) {
		handleResponse(w, http.StatusBadRequest, nil)
		return
	}

	pKey, err := h.storage.Customer().Create(createCustomers)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	customer, err := h.storage.Customer().Get(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, customer)
}

//getbyidcustomers
func (h Handler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	customers, err := h.storage.Customer().Get(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(customers)
	handleResponse(w, http.StatusOK, customers)
}

//getall customers
func (h Handler) GetCustomerList(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.storage.Customer().GetList(models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

//update customers
func (h Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	updateCustomers := models.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&updateCustomers); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Customer().Update(updateCustomers)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	customer, err := h.storage.City().Get(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, customer)
}

//delete customers
func (h Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	if err := h.storage.Customer().Delete(models.PrimaryKey{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "data successfully deleted")
}
