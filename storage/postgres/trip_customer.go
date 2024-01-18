package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	uid := uuid.New()
	createdAt := time.Now()

	query := `
        INSERT INTO trip_customers (id, trip_id, customer_id, created_at)
        VALUES ($1, $2, $3, $4)
    `

	_, err := c.db.Exec(query, uid, req.TripID, req.CustomerID, createdAt)
	if err != nil {
		fmt.Println("error while creating trip customer:", err.Error())
		return "", err
	}

	return uid.String(), nil
}

func (c *tripCustomerRepo) Get(id models.PrimaryKey) (models.TripCustomer, error) {
	tripcustomer := models.TripCustomer{}

	query := `
        SELECT id, trip_id, customer_id, created_at FROM trip_customers
        WHERE id = $1
    `

	if err := c.db.QueryRow(query, id.ID).Scan(
		&tripcustomer.ID,
		&tripcustomer.TripID,
		&tripcustomer.CustomerID,
		&tripcustomer.CreatedAt,
	); err != nil {
		fmt.Println("error while scanning trip", err.Error())
		return models.TripCustomer{}, err
	}

	return models.TripCustomer{}, nil
}

func (c *tripCustomerRepo) GetList(req models.GetListRequest) (models.TripCustomersResponse, error) {
	var (
		tripsCustomers    = []models.TripCustomer{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
		SELECT count(1) from trip_customers  `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of tripscustomers", err.Error())
		return models.TripCustomersResponse{}, err
	}

	query = `
		SELECT id, trip_id, customer_id, created_at
			FROM trip_customers
			   
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.TripCustomersResponse{}, err
	}

	for rows.Next() {
		tripcustomer := models.TripCustomer{}

		if err = rows.Scan(
			&tripcustomer.ID,
			&tripcustomer.TripID,
			&tripcustomer.CustomerID,
			&tripcustomer.CreatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.TripCustomersResponse{}, err
		}

		tripsCustomers = append(tripsCustomers, tripcustomer)
	}

	return models.TripCustomersResponse{
		TripCustomers: tripsCustomers,
		Count:         count,
	}, nil

}

func (c *tripCustomerRepo) Update(req models.TripCustomer) (string, error) {
	return "", nil
}

func (c *tripCustomerRepo) Delete(id models.PrimaryKey) error {
	query := `
	delete from trip_customers
	WHERE id = $1
`
	if _, err := c.db.Exec(query,id.ID); err != nil {
		fmt.Println("error while deleting trip_customers by id", err.Error())
		return err
	}

	return nil
}
