package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
)

type tripRepo struct {
	db *sql.DB
}

func NewTripRepo(db *sql.DB) storage.ITripRepo {
	return &tripRepo{
		db: db,
	}
}

func (c *tripRepo) Create(req models.CreateTrip) (string, error) {
	return "", nil
}

func (c *tripRepo) Get(id string) (models.Trip, error) {
	return models.Trip{}, nil
}

func (c *tripRepo) GetList(req models.GetListRequest) (models.TripsResponse, error) {
	return models.TripsResponse{}, nil
}

func (c *tripRepo) Update(req models.Trip) (string, error) {
	return "", nil
}

func (c *tripRepo) Delete(id string) error {
	return nil
}
