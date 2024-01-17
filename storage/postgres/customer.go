package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
)

type customerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) storage.ICustomerRepo {
	return customerRepo{
		db,
	}
}

func (c customerRepo) Create(car models.CreateCustomer) (string, error) {

	return "", nil
}

func (c customerRepo) Get(id string) (models.Customer, error) {
	return models.Customer{}, nil
}

func (c customerRepo) GetList(req models.GetListRequest) (models.CustomersResponse, error) {

	return models.CustomersResponse{}, nil
}

func (c customerRepo) Update(car models.Customer) (string, error) {

	return "", nil
}

func (c customerRepo) Delete(id string) error {

	return nil
}
