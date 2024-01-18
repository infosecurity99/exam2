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

func (h Handler) City(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCity(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetCityList(w, r)
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

	newsCity := models.CreateCity{}

	if err := json.NewDecoder(r.Body).Decode(&newsCity); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cityID, err := h.storage.City().Create(newsCity)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if check.NameCity(models.CreateCity{}) {
		fmt.Println("Invalid name length")
		handleResponse(w, http.StatusBadRequest, "Invalid name length")
		return
	}

	createdCity, err := h.storage.City().Get((models.PrimaryKey{ID: cityID}).ID)

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, createdCity)
}

//getcitybyid
func (h Handler) GetCityByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	user, err := h.storage.City().Get(models.PrimaryKey{
		ID: id,
	}.ID)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(user)
	handleResponse(w, http.StatusOK, user)
}

func (h Handler) GetCityList(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.storage.City().GetList(models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

//update city
func (h Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	updateCity := models.City{}

	if err := json.NewDecoder(r.Body).Decode(&updateCity); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.City().Update(updateCity)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storage.City().Get(models.PrimaryKey{
		ID: pKey,
	}.ID)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, user)
}

//delete city
func (h Handler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	if err := h.storage.City().Delete(models.PrimaryKey{ID: id}.ID); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "data successfully deleted")
}
