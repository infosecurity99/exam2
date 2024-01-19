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

//create  insert
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

//getbyid
func (c *tripCustomerRepo) Get(id models.PrimaryKey) (models.TripCustomer, error) {
	tripcustomer := models.TripCustomer{}
	customer := models.Customer{} // Assuming you have a Customer struct

	query := `
        SELECT
            tc.id,
            tc.trip_id,
            tc.customer_id,
            tc.created_at,
            c.id AS customer_id,
            c.full_name AS customer_name,
            c.email AS customer_email,
            c.phone AS customer_phone
        FROM
            trip_customers tc
        JOIN
            customers c ON tc.customer_id = c.id
        WHERE
            tc.id = $1
    `

	if err := c.db.QueryRow(query, id.ID).Scan(
		&tripcustomer.ID,
		&tripcustomer.TripID,
		&tripcustomer.CustomerID,
		&tripcustomer.CreatedAt,
		&customer.ID,
		&customer.FullName,
		&customer.Email,
		&customer.Phone,
	); err != nil {
		fmt.Println("error while scanning trip", err.Error())
		return models.TripCustomer{}, err
	}

	tripcustomer.CustomerData = customer

	return tripcustomer, nil
}

//get list trip customers
func (c *tripCustomerRepo) GetList(req models.GetListRequest) (models.TripCustomersResponse, error) {
	var (
		tripsCustomers    = []models.TripCustomer{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
        SELECT count(1) FROM trip_customers  
    `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of trip customers", err.Error())
		return models.TripCustomersResponse{}, err
	}

	query = `
        SELECT
            tc.id,
            tc.trip_id,
            tc.customer_id,
            tc.created_at,
            c.id AS customer_id,
            c.full_name AS customer_name,
            c.email AS customer_email,
            c.phone AS customer_phone
        FROM
            trip_customers tc
        JOIN
            customers c ON tc.customer_id = c.id
        LIMIT $1 OFFSET $2
    `

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while querying rows", err.Error())
		return models.TripCustomersResponse{}, err
	}

	for rows.Next() {
		tripcustomer := models.TripCustomer{}
		customer := models.Customer{}

		if err = rows.Scan(
			&tripcustomer.ID,
			&tripcustomer.TripID,
			&tripcustomer.CustomerID,
			&tripcustomer.CreatedAt,
			&customer.ID,
			&customer.FullName,
			&customer.Email,
			&customer.Phone,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.TripCustomersResponse{}, err
		}

		// Assuming you have a Customer field in your TripCustomer struct
		tripcustomer.CustomerData = customer

		tripsCustomers = append(tripsCustomers, tripcustomer)
	}

	return models.TripCustomersResponse{
		TripCustomers: tripsCustomers,
		Count:         count,
	}, nil
}

func (c *tripCustomerRepo) Update(req models.TripCustomer) (string, error) {
	query := `
        UPDATE trip_customers 
        SET
            trip_id = $1,
            customer_id = $2

        WHERE id = $3
    `

	_, err := c.db.Exec(query, req.TripID, req.CustomerID, req.ID)
	if err != nil {
		fmt.Println("error while updating trip customer data:", err.Error())
		return "", err
	}

	return req.ID, nil
}

func (c *tripCustomerRepo) Delete(id models.PrimaryKey) error {
	query := `
	delete from trip_customers
	WHERE id = $1
`
	if _, err := c.db.Exec(query, id.ID); err != nil {
		fmt.Println("error while deleting trip_customers by id", err.Error())
		return err
	}

	return nil
}
