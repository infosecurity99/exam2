package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) storage.IDriverRepo {
	return driverRepo{
		DB: db,
	}
}

//  create drivers
func (d driverRepo) Create(driver models.CreateDriver) (string, error) {
	uid := uuid.New()
	createAt := time.Now()

	_, err := d.DB.Exec(`
		insert into drivers (id, full_name, phone, from_city_id, to_city_id, created_at)
		values ($1, $2, $3, $4, $5, $6)
	`, uid, driver.FullName, driver.Phone, driver.FromCityID, driver.ToCityID, createAt)

	if err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyid drivers
func (d driverRepo) Get(id string) (models.Driver, error) {
	driver := models.Driver{}

	query := `
		SELECT id, full_name, phone, from_city_id, to_city_id, created_at FROM drivers WHERE id = $1
	`

	err := d.DB.QueryRow(query, id).Scan(
		&driver.ID,
		&driver.FullName,
		&driver.Phone,
		&driver.FromCityID,
		&driver.ToCityID,
		&driver.CreatedAt,
	)

	if err != nil {
		fmt.Println("error while scanning driver", err.Error())
		return models.Driver{}, err
	}

	return driver, nil
}

//getlist drivers
func (d driverRepo) GetList(req models.GetListRequest) (models.DriversResponse, error) {
	var (
		drivers         = []models.Driver{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
		SELECT count(1) from drivers  `

	if err := d.DB.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of drivers", err.Error())
		return models.DriversResponse{}, err
	}

	query = `
		SELECT id, full_name,phone,email, created_at
			FROM customers
			   
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := d.DB.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.DriversResponse{}, err
	}

	for rows.Next() {
		costumer := models.Customer{}

		if err = rows.Scan(
			&costumer.ID,
			&costumer.FullName,
			&costumer.Phone,
			&costumer.Email,
			&costumer.CreatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.DriversResponse{}, err
		}

		costumers = append(costumers, costumer)
	}

	return models.DriversResponse{
		Customers: costumers,
		Count:     count,
	}, nil
}

//update  drivers
func (d driverRepo) Update(driver models.Driver) (string, error) {

	return "", nil
}

//delete drivers
func (d driverRepo) Delete(id string) error {

	return nil
}
