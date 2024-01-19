package handler

import (
	"encoding/json"
	"errors"
	"exam2/api/models"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Trip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTrip(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetTripList(w, r)
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

	newsTrip := models.CreateTrip{}

	if err := json.NewDecoder(r.Body).Decode(&newsTrip); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tripID, err := h.storage.Trip().Create(newsTrip)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdTrip, err := h.storage.Trip().Get(models.PrimaryKey{ID: tripID})

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusCreated, createdTrip)
}

//getbyidtrip
func (h Handler) GetTripByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	trip, err := h.storage.Trip().Get(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, trip)
}

//getlist trip
func (h Handler) GetTripList(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.storage.Trip().GetList(models.GetListRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

//updatetrip
func (h Handler) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	updateTrip := models.Trip{}

	if err := json.NewDecoder(r.Body).Decode(&updateTrip); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Trip().Update(updateTrip)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storage.Trip().Get(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, user)
}

//deleate trip
func (h Handler) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	if err := h.storage.Trip().Delete(models.PrimaryKey{ID: id}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "data successfully deleted")
}
