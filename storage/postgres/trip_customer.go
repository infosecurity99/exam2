package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
)

type tripCustomerRepo struct {
	db *sql.DB
}

func NewTripCustomerRepo(db *sql.DB) storage.ITripCustomerRepo {
	return &tripCustomerRepo{
		db: db,
	}
}

func (c *tripCustomerRepo) Create(req models.CreateTripCustomer) (string, error) {
	return "", nil
}

func (c *tripCustomerRepo) Get(id string) (models.TripCustomer, error) {
	return models.TripCustomer{}, nil
}

func (c *tripCustomerRepo) GetList(req models.GetListRequest) (models.TripCustomersResponse, error) {
	return models.TripCustomersResponse{}, nil
}

func (c *tripCustomerRepo) Update(req models.TripCustomer) (string, error) {
	return "", nil
}

func (c *tripCustomerRepo) Delete(id string) error {
	return nil
}
