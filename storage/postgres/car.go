package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type carRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) storage.ICarRepo {
	return carRepo{
		db: db,
	}
}

//create car
func (c carRepo) Create(car models.CreateCar) (string, error) {
	uid := uuid.New()
	createAt := time.Now()

	if _, err := c.db.Exec(`
        INSERT INTO cars (id, model, brand, number,driver_id, created_at)
        VALUES ($1, $2, $3, $4, $5, $6 )`,
		uid,
		car.Model,
		car.Brand,
		car.Number,
		car.DriverID,
		createAt,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyidcar
func (c carRepo) Get(pkey models.PrimaryKey) (models.Car, error) {
	car := models.Car{}
	var driverData models.Driver

	err := c.db.QueryRow(`
        SELECT
            cars.id,
            cars.model,
            cars.brand,
            cars.number,
         
            cars.driver_id,
            drivers_from.id AS driver_id,
            drivers_from.full_name AS driver_full_name,
            drivers_from.phone AS driver_phone,
            cars.created_at
        FROM
            cars
        LEFT JOIN 
            drivers AS drivers_from ON cars.driver_id = drivers_from.id
        WHERE 
            cars.id = $1
    `, pkey.ID).Scan(
		&car.ID,
		&car.Model,
		&car.Brand,
		&car.Number,

		&car.DriverID,
		&driverData.ID,
		&driverData.FullName,
		&driverData.Phone,
		&car.CreatedAt,
	)

	if err != nil {
		log.Printf("Error while retrieving car data. Error: %s\n", err.Error())
		return models.Car{}, err
	}

	car.DriverData = driverData
	return car, nil
}

// getlistcar with driver data
func (c carRepo) GetList(request models.GetListRequest) (models.CarsResponse, error) {
	var (
		cars  = []models.Car{}
		count = 0
		query string
	)

	countQuery := `
        select count(1) from cars `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of customers", err.Error())
		return models.CarsResponse{}, err
	}

	var offset int
	offset = (request.Page - 1) * request.Limit
	if offset < 0 {
		offset = 0
	}

	query = `
        SELECT
            cars.id,
            cars.model,
            cars.brand,
            cars.number,
      
            cars.driver_id,
            drivers.id AS driver_id,
            drivers.full_name AS driver_full_name,
            drivers.phone AS driver_phone,
            drivers.from_city_id AS from_city_id,
            drivers.to_city_id AS to_city_id,
            from_city.id AS from_city_id,
            from_city.name AS from_city_name,
            from_city.created_at AS from_city_created_at,
			drivers.created_at AS created_at,
            to_city.id AS to_city_id,
            to_city.name AS to_city_name,
            to_city.created_at AS to_city_created_at,
			drivers.created_at AS created_at,
            cars.created_at
        FROM
            cars
        LEFT JOIN 
            drivers ON cars.driver_id = drivers.id
        LEFT JOIN
            cities AS from_city ON drivers.from_city_id = from_city.id
        LEFT JOIN
            cities AS to_city ON drivers.to_city_id = to_city.id
        `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, request.Limit, offset)
	if err != nil {
		log.Printf("Error while querying rows. Error: %s\n", err.Error())
		return models.CarsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car

		if err = rows.Scan(
			&car.ID,
			&car.Model,
			&car.Brand,
			&car.Number,

			&car.DriverID,
			&car.DriverData.ID,
			&car.DriverData.FullName,
			&car.DriverData.Phone,
			&car.DriverData.FromCityID,
			&car.DriverData.ToCityID,
			&car.DriverData.FromCityData.ID,
			&car.DriverData.FromCityData.Name,
			&car.DriverData.FromCityData.CreatedAt,
			&car.DriverData.CreatedAt,
			&car.DriverData.ToCityData.ID,
			&car.DriverData.ToCityData.Name,
			&car.DriverData.ToCityData.CreatedAt,
			&car.DriverData.CreatedAt,
			&car.CreatedAt,
		); err != nil {
			log.Println("Error while scanning row:", err)
			return models.CarsResponse{}, err
		}

		cars = append(cars, car)
	}

	return models.CarsResponse{
		Cars:  cars,
		Count: count,
	}, nil
}

//updatecar
func (c carRepo) Update(car models.Car) (string, error) {
	query := `
	UPDATE cars 
	SET model = $1,brand=$2, number=$3, created_at = $4
	WHERE id = $5
`

	_, err := c.db.Exec(query, car.Model, car.Brand, car.Number, car.CreatedAt, car.ID)
	if err != nil {
		fmt.Println("error while updating cars data:", err.Error())
		return " ", err
	}

	return car.ID, nil
}

//delete car
func (c carRepo) Delete(i models.PrimaryKey) error {
	query := `
	delete from cars
		where id = $1
`
	if _, err := c.db.Exec(query, i.ID); err != nil {
		fmt.Println("error while deleting cars by id", err.Error())
		return err
	}

	return nil
}

func (c carRepo) UpdateCarRoute(updateCarRoute models.UpdateCarRoute) error {

	return nil
}

func (c carRepo) UpdateCarStatus(updateCarStatus models.UpdateCarStatus) error {

	return nil
}
