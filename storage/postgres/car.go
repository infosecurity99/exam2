package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type carRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) storage.ICarRepo {
	return carRepo{
		db,
	}
}

//create car
func (c carRepo) Create(car models.CreateCar) (string, error) {
	uid := uuid.New()
	createAt := time.Now()

	if _, err := c.db.Exec(`
        INSERT INTO cars (id, model, brand, number, status,driver_id, created_at)
        VALUES ($1, $2, $3, $4, $5, $6 , $7)`,
		uid,
		car.Model,
		car.Brand,
		car.Number,
		car.Status,
		car.DriverID,
		createAt,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyidcar
func (c carRepo) Get(id string) (models.Car, error) {
	car := models.Car{}

	query := `
		SELECT id, model, brand, number, created_at FROM cars WHERE id = $1
	`

	if err := c.db.QueryRow(query, id).Scan(
		&car.ID,
		&car.Model,
		&car.Brand,
		&car.Number,

		&car.CreatedAt,
	); err != nil {
		fmt.Println("error while scanning car", err.Error())
		return models.Car{}, err
	}

	return car, nil
}

//getlistcar
// getlistcar with driver data
func (c carRepo) GetList(req models.GetListRequest) (models.CarsResponse, error) {
	var (
		cars              = []models.Car{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
		SELECT count(1) from cars `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of cars", err.Error())
		return models.CarsResponse{}, err
	}

	query = `
		SELECT c.id, c.model, c.brand, c.number, c.created_at, d.full_name as driver_name, d.phone as driver_phone
			FROM cars c
			JOIN drivers d ON c.driver_id = d.id
    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.CarsResponse{}, err
	}

	for rows.Next() {
		car := models.Car{}
		driverName, driverPhone := "", ""

		if err = rows.Scan(
			&car.ID,
			&car.Model,
			&car.Brand,
			&car.Number,
			&car.CreatedAt,
			&driverName,
			&driverPhone,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
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
func (c carRepo) Delete(id string) error {
	query := `
	delete from cars
		where id = $1
`
	if _, err := c.db.Exec(query, models.PrimaryKey{ID: id}.ID); err != nil {
		fmt.Println("error while deleting cars by id", err.Error())
		return err
	}

	return nil
}

func (c carRepo) UpdateCarRoute(updateCarRoute models.UpdateCarRoute) error {
	query := `
        UPDATE cars 
        SET from_city = $1, to_city = $2
        WHERE id = $3
    `

	_, err := c.db.Exec(query, updateCarRoute.FromCityID, updateCarRoute.ToCityID, updateCarRoute.CarID)
	if err != nil {
		fmt.Println("error while updating car route:", err.Error())
		return err
	}

	return nil
}

func (c carRepo) UpdateCarStatus(updateCarStatus models.UpdateCarStatus) error {
	query := `
        UPDATE cars 
        SET status = $1
        WHERE id = $2
    `

	_, err := c.db.Exec(query, updateCarStatus.Status, updateCarStatus.ID)
	if err != nil {
		fmt.Println("error while updating car status:", err.Error())
		return err
	}

	return nil
}
