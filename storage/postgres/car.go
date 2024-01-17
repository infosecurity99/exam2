package postgres

import (
	"exam2/storage"
	"database/sql"
	"exam2/api/models"
)

type carRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) storage.ICarRepo {
	return carRepo{
		db,
	}
}

func (c carRepo) Create(car models.CreateCar) (string, error) {

	return "", nil
}

func (c carRepo) Get(id string) (models.Car, error) {
	return models.Car{}, nil
}

func (c carRepo) GetList(req models.GetListRequest) (models.CarsResponse, error) {

	return models.CarsResponse{}, nil
}

func (c carRepo) Update(car models.Car) (string, error) {
	return "", nil
}

func (c carRepo) Delete(id string) error {

	return nil
}

func (c carRepo) UpdateCarRoute(models.UpdateCarRoute) error {

	return nil
}
func (c carRepo) UpdateCarStatus(updateCarStatus models.UpdateCarStatus) error {

	return nil
}
