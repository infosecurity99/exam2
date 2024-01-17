package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type cityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) storage.ICityRepo {
	return cityRepo{
		db,
	}
}

//  create city
func (c cityRepo) Create(city models.CreateCity) (string, error) {
	uid := uuid.New()
	createat := time.Now()

	if _, err := c.db.Exec(`insert into cities (id, name, created_at) values ($1, $2, $3)`,
		uid,
		city.Name,
		createat,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
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
