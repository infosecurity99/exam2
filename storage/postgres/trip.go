package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
)

type tripRepo struct {
	db *sql.DB
}

func NewTripRepo(db *sql.DB) storage.ITripRepo {
	return &tripRepo{
		db: db,
	}
}
//create trip
func (c *tripRepo) Create(req models.CreateTrip) (string, error) {
	return "", nil
}
//getbyidtrip
func (c *tripRepo) Get(id string) (models.Trip, error) {
	return models.Trip{}, nil
}
//getlisttrip
func (c *tripRepo) GetList(req models.GetListRequest) (models.TripsResponse, error) {
	return models.TripsResponse{}, nil
}
//updatetrip
func (c *tripRepo) Update(req models.Trip) (string, error) {
	return "", nil
}
//delete trip
func (c *tripRepo) Delete(id string) error {
	return nil
}
