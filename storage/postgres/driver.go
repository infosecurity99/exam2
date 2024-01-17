package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
)

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) storage.IDriverRepo {
	return driverRepo{
		DB: db,
	}
}

func (d driverRepo) Create(driver models.CreateDriver) (string, error) {
	return "", nil
}

func (d driverRepo) Get(id string) (models.Driver, error) {
	return models.Driver{}, nil
}

func (d driverRepo) GetList(req models.GetListRequest) (models.DriversResponse, error) {

	return models.DriversResponse{}, nil
}

func (d driverRepo) Update(driver models.Driver) (string, error) {

	return "", nil
}

func (d driverRepo) Delete(id string) error {

	return nil
}
