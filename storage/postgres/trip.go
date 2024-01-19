package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type tripRepo struct {
	db *sql.DB
}

func NewTripRepo(db *sql.DB) storage.ITripRepo {
	return &tripRepo{
		db: db,
	}
}

//create trip
func (c *tripRepo) Create(req models.CreateTrip) (string, error) {
	uid := uuid.New()
	createdAt := time.Now()

	if _, err := c.db.Exec(`
       insert into trips (id, trip_number_id, from_city_id, to_city_id, driver_id, price, created_at)
        values ($1, $2, $3, $4, $5, $6, $7)`,
		uid,
		req.TripNumberID,
		req.FromCityID,
		req.ToCityID,
		req.DriverID,
		req.Price,
		createdAt,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//getbyidtrip
func (c tripRepo) Get(id models.PrimaryKey) (models.Trip, error) {
	trip := models.Trip{}

	query := `
        SELECT id, trip_number_id, from_city_id, to_city_id, driver_id, price, created_at FROM trips
        WHERE id = $1
    `

	if err := c.db.QueryRow(query, id.ID).Scan(
		&trip.ID,
		&trip.TripNumberID,
		&trip.FromCityID,
		&trip.ToCityID,
		&trip.DriverID,
		&trip.Price,
		&trip.CreatedAt,
	); err != nil {
		fmt.Println("error while scanning trip", err.Error())
		return models.Trip{}, err
	}

	return trip, nil
}

//getlisttrip
func (c tripRepo) GetList(req models.GetListRequest) (models.TripsResponse, error) {
	var (
		trips             = []models.Trip{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
		SELECT count(1) from trips  `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of trips", err.Error())
		return models.TripsResponse{}, err
	}

	query = `
		SELECT id, trip_number_id, from_city_id, to_city_id, driver_id, price, created_at
			FROM trips
			   
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.TripsResponse{}, err
	}

	for rows.Next() {
		trip := models.Trip{}

		if err = rows.Scan(
			&trip.ID,
			&trip.TripNumberID,
			&trip.FromCityID,
			&trip.ToCityID,
			&trip.DriverID,
			&trip.Price,
			&trip.CreatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.TripsResponse{}, err
		}

		trips = append(trips, trip)
	}

	return models.TripsResponse{
		Trips: trips,
		Count: count,
	}, nil
}

//updatetrip
func (c tripRepo) Update(req models.Trip) (string, error) {
	query := `
        UPDATE trips 
        SET trip_number_id = $1, 
            from_city_id = $2, 
            to_city_id = $3, 
            driver_id = $4, 
            price = $5, 
            created_at = $6
        WHERE id = $7
    `

	_, err := c.db.Exec(query, req.TripNumberID, req.FromCityID, req.ToCityID, req.DriverID, req.Price, req.CreatedAt, req.ID)
	if err != nil {
		fmt.Println("error while updating trips data:", err.Error())
		return " ", err
	}

	return req.ID, nil
}

//delete trip
func (c tripRepo) Delete(id models.PrimaryKey) error {
	query := `
        delete from trips
        WHERE id = $1
    `
	if _, err := c.db.Exec(query, id.ID); err != nil {
		fmt.Println("error while deleting trip by id", err.Error())
		return err
	}

	return nil
}
