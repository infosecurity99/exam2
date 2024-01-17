package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
)

type cityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) storage.ICityRepo {
	return cityRepo{
		db,
	}
}

func (c cityRepo) Create(car models.CreateCity) (string, error) {

	return "", nil
}

func (c cityRepo) Get(id string) (models.City, error) {
	return models.City{}, nil
}

func (c cityRepo) GetList(req models.GetListRequest) (models.CitiesResponse, error) {

	return models.CitiesResponse{}, nil
}

func (c cityRepo) Update(car models.City) (string, error) {

	return "", nil
}

func (c cityRepo) Delete(id string) error {

	return nil
}
